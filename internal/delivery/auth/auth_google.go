package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	modelsAuth "forum/internal/models/auth"
	modelsData "forum/internal/models/data"
	modelsUser "forum/internal/models/user"
	serviceAuth "forum/internal/service/auth"
	"forum/pkg"
	"log"
	"net/http"
	"net/url"
	"strings"
	"text/template"
)

func requestToGoogle(w http.ResponseWriter, r *http.Request, cfg *modelsAuth.OAuthConfig) {
	URL, err := url.Parse(googleAuthURL)
	if err != nil {
		log.Printf("Parse: %s", err)
	}

	parameters := url.Values{}
	parameters.Add("client_id", cfg.ClientID)
	parameters.Add("redirect_uri", cfg.RedirectURL)
	parameters.Add("scope", strings.Join(cfg.Scopes, " "))
	parameters.Add("response_type", "code")
	URL.RawQuery = parameters.Encode()
	url := URL.String()
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (a *AuthHandler) GoogleSignIn(w http.ResponseWriter, r *http.Request) {
	requestToGoogle(w, r, googleSignInConfig)
}

func (a *AuthHandler) GoogleSignUp(w http.ResponseWriter, r *http.Request) {
	requestToGoogle(w, r, googleSignUpConfig)
}

func (a *AuthHandler) SignInCallbackFromGoogle(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")

	user, err := userFromGoogleInfo(code, googleSignInConfig)
	// fmt.Println(user)
	if err != nil {
		pkg.Error(w, http.StatusUnauthorized, err)
		return
	}
	if err := a.SetSession(w, user); err != nil {
		if errors.Is(err, serviceAuth.ErrNoUser) || errors.Is(err, serviceAuth.ErrWrongPassword) {
			pkg.Error(w, http.StatusUnauthorized, err)
			return
		}
		pkg.Error(w, http.StatusInternalServerError, err)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (a *AuthHandler) SignUpCallbackFromGoogle(w http.ResponseWriter, r *http.Request) {
	var tmplData modelsData.Data
	switch r.Method {
	case http.MethodGet:
		tmpl, err := template.ParseFiles("./ui/templates/oauth2.html")
		if err != nil {
			pkg.ErrorLog.Println(err)
			pkg.Error(w, http.StatusInternalServerError, err)
			return
		}
		tmpl.Execute(w, tmplData)
	case http.MethodPost:
		// var userForSign modelsUser.User
		code := r.URL.Query().Get("code")
		user, err := userFromGoogleInfo(code, googleSignUpConfig)
		if err != nil {

			pkg.Error(w, http.StatusUnauthorized, err)
			return
		}

		if err := r.ParseForm(); err != nil {
			pkg.Error(w, http.StatusInternalServerError, err)
			return
		}
		username := r.Form["username"]

		user.Username = username[0]
		if err := a.user.CreateUser(*user); err != nil {
			if errors.Is(err, serviceAuth.ErrEmailTaken) || errors.Is(err, serviceAuth.ErrUsernameTaken) {
				pkg.Error(w, http.StatusBadRequest, err)
				return
			}
			pkg.Error(w, http.StatusInternalServerError, err)
			return
		}

		if err := a.SetSession(w, user); err != nil {
			if errors.Is(err, serviceAuth.ErrNoUser) {
				fmt.Println("owibka")
				pkg.Error(w, http.StatusUnauthorized, err)
				return
			}
			pkg.Error(w, http.StatusInternalServerError, err)
		}
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func userFromGoogleInfo(code string, cfg *modelsAuth.OAuthConfig) (*modelsUser.User, error) {
	accessToken, err := googleAccessToken(cfg, code)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		"GET",
		"https://www.googleapis.com/oauth2/v2/userinfo?access_token="+url.QueryEscape(accessToken),
		nil,
	)
	if err != nil {
		return nil, err
	}
	// fmt.Println(req)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var u *modelsAuth.UserInfo
	if err := json.NewDecoder(resp.Body).Decode(&u); err != nil {
		return nil, err
	}
	// fmt.Println(json.NewDecoder(resp.Body).Decode(&u))

	if u.Email == "" {
		return nil, errors.New("email is empty")
	}

	user := &modelsUser.User{
		Email:      u.Email,
		AuthMethod: "google",
	}

	return user, nil
}

func googleAccessToken(cfg *modelsAuth.OAuthConfig, code string) (string, error) {
	v := url.Values{
		"grant_type":    {"authorization_code"},
		"code":          {code},
		"redirect_uri":  {cfg.RedirectURL},
		"client_id":     {cfg.ClientID},
		"client_secret": {cfg.ClientSecret},
	}
	req, err := http.NewRequest("POST", googleTokenURL, strings.NewReader(v.Encode()))
	if err != nil {
		return "", err
	}
	// fmt.Println(&req.Body, "reqBody")

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	var token modelsAuth.Token
	if err := json.NewDecoder(resp.Body).Decode(&token); err != nil {
		return "", err
	}

	return token.AccessToken, nil
}
