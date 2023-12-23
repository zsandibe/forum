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

func requestToGithub(w http.ResponseWriter, r *http.Request, cfg modelsAuth.OAuthConfig) {
	URL, err := url.Parse(githubAuthURL)
	if err != nil {
		log.Printf("Parse: %s", err)
	}

	parameters := url.Values{}
	parameters.Add("client_id", cfg.ClientID)
	parameters.Add("redirect_uri", cfg.RedirectURL)
	parameters.Add("scope", strings.Join(cfg.Scopes, " "))

	URL.RawQuery = parameters.Encode()
	url := URL.String()
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (a *AuthHandler) GithubSignIn(w http.ResponseWriter, r *http.Request) {
	requestToGithub(w, r, *githubSignInConfig)
}

func (a *AuthHandler) GithubSignUp(w http.ResponseWriter, r *http.Request) {
	requestToGithub(w, r, *githubSignUpConfig)
}

func (a *AuthHandler) SignInCallbackFromGithub(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")

	user, err := userFromGithubInfo(code, githubSignInConfig)
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

func (a *AuthHandler) SignUpCallbackFromGithub(w http.ResponseWriter, r *http.Request) {
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
		code := r.URL.Query().Get("code")

		user, err := userFromGithubInfo(code, githubSignUpConfig)
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
				pkg.Error(w, http.StatusUnauthorized, err)
				return
			}
			pkg.Error(w, http.StatusInternalServerError, err)
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	default:
		pkg.Error(w, http.StatusMethodNotAllowed, nil)
	}
}

func userFromGithubInfo(code string, cfg *modelsAuth.OAuthConfig) (*modelsUser.User, error) {
	accessToken, err := githubAccessToken(cfg, code)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		"GET",
		"https://api.github.com/user/emails",
		nil,
	)
	if err != nil {
		return nil, err
	}

	authHeaderValue := fmt.Sprintf("token %s", accessToken)
	req.Header.Set("Authorization", authHeaderValue)
	req.Header.Set("accept", "application/vnd.github.v3+json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var emails []*modelsAuth.UserInfo
	if err := json.NewDecoder(resp.Body).Decode(&emails); err != nil {
		return nil, err
	}

	if emails[0].Email == "" {
		return nil, errors.New("email is empty")
	}

	user := &modelsUser.User{
		Email:      emails[0].Email,
		AuthMethod: "github",
	}

	return user, nil
}

func githubAccessToken(cfg *modelsAuth.OAuthConfig, code string) (string, error) {
	v := url.Values{
		"code":          {code},
		"client_id":     {cfg.ClientID},
		"client_secret": {cfg.ClientSecret},
	}

	req, err := http.NewRequest("POST", githubTokenURL, strings.NewReader(v.Encode()))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")

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
