package auth

import (
	"errors"
	"fmt"
	"forum/pkg"
	"net/http"
	"text/template"

	modelsData "forum/internal/models/data"
	modelsUser "forum/internal/models/user"
	serviceAuth "forum/internal/service/auth"
)

func (a *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("ok")
	// fmt.Println("OK")s
	var tmplData modelsData.Data
	if r.Method == http.MethodGet {
		tmpl, err := template.ParseFiles("./ui/templates/sign-up.html")
		if err != nil {
			pkg.ErrorLog.Println(err)
			pkg.Error(w, http.StatusInternalServerError, err)
			return
		}
		tmpl.Execute(w, tmplData)
	} else if r.Method == http.MethodPost {
		// fmt.Println("OK")
		if err := r.ParseForm(); err != nil {
			pkg.ErrorLog.Println(err)
			pkg.Error(w, http.StatusInternalServerError, err)
			return
		}
		username, flag1 := r.Form["username"]
		email, flag2 := r.Form["email"]
		password, flag3 := r.Form["password"]
		confirmedPassword, flag4 := r.Form["confirm-password"]
		// fmt.Println(email, username, password, confirmedPassword)
		if !flag1 || !flag2 || !flag3 || !flag4 {
			fmt.Println("OK ERROR")
			pkg.Error(w, http.StatusBadRequest, nil)
			return
		}
		if password[0] != confirmedPassword[0] {
			pkg.Error(w, http.StatusBadRequest, errors.New("Password and confirmed password do not match"))
			return
		}
		role := "user"
		user := modelsUser.User{
			Email:           email[0],
			Username:        username[0],
			Password:        password[0],
			ConfirmPassword: confirmedPassword[0],
			Role:            role,
		}
		if err := a.user.CreateUser(user); err != nil {
			if errors.Is(err, pkg.ErrInvalidEmail) || errors.Is(err, pkg.ErrInvalidPassword) ||
				errors.Is(err, pkg.ErrInvalidUsername) || errors.Is(err, serviceAuth.ErrUsernameTaken) ||
				errors.Is(err, serviceAuth.ErrEmailTaken) {
				// fmt.Println("OK")
				pkg.Error(w, http.StatusBadRequest, err)
				return
			}
			pkg.Error(w, http.StatusInternalServerError, err)
			return
		}
		http.Redirect(w, r, "/auth/sign-in", http.StatusSeeOther)
	} else {
		pkg.Error(w, http.StatusMethodNotAllowed, nil)
		// fmt.Println("OK")
		return
	}
}

func (a *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	var tmplData modelsData.Data
	if r.Method == http.MethodGet {
		tmpl, err := template.ParseFiles("./ui/templates/sign-in.html")
		if err != nil {
			pkg.ErrorLog.Println(err)
			pkg.Error(w, http.StatusInternalServerError, err)
			return
		}
		tmpl.Execute(w, tmplData)
	} else if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			pkg.ErrorLog.Println(err)
			pkg.Error(w, http.StatusInternalServerError, err)
			return
		}
		email, ok1 := r.Form["email"]
		password, ok2 := r.Form["password"]
		if !ok1 || !ok2 {
			pkg.ErrorLog.Println(ok1, ok2)
			pkg.Error(w, http.StatusBadRequest, nil)
			return
		}
		user := modelsUser.User{
			Email:    email[0],
			Password: password[0],
		}
		if err := a.SetSession(w, &user); err != nil {
			// fmt.Println(err)
			// fmt.Println("error setting session")
			if errors.Is(err, serviceAuth.ErrNoUser) || errors.Is(err, serviceAuth.ErrWrongPassword) {
				pkg.Error(w, http.StatusUnauthorized, err)
				return
			}

			pkg.Error(w, http.StatusInternalServerError, err)

		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		pkg.Error(w, http.StatusMethodNotAllowed, nil)
		return
	}
}

func (a *AuthHandler) LogOut(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		pkg.Error(w, http.StatusMethodNotAllowed, nil)
		return
	}
	cookie, err := r.Cookie("session")
	if err != nil {
		pkg.ErrorLog.Println(err)
		http.Redirect(w, r, "/auth/sign-in", http.StatusSeeOther)
		return
	}
	if err := a.user.DeleteSession(cookie.Value); err != nil {
		pkg.ErrorLog.Println(err)
		pkg.Error(w, http.StatusInternalServerError, err)
		return
	}
	// http.SetCookie(w, &http.Cookie{
	// 	Name:     "session",
	// 	Value:    "",
	// 	Expires:  time.Now(),
	// 	HttpOnly: false,
	// 	MaxAge:   -1,
	// })
	cookie.MaxAge = -1
	cookie.Name = "session"
	cookie.Value = ""
	cookie.Path = "/"
	cookie.HttpOnly = false
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (a *AuthHandler) GetAllUsersList(w http.ResponseWriter, r *http.Request) {

	userValue := r.Context().Value(modelsData.ContextKeyUser)
	if userValue == nil {
		pkg.Error(w, http.StatusUnauthorized, nil)
		return
	}
	user, ok := userValue.(modelsUser.User)
	if !ok {
		pkg.Error(w, http.StatusInternalServerError, nil)
		return
	}
	if r.Method == http.MethodGet {
		users, err := a.user.GetAllUsersList()
		if err != nil {
			pkg.Error(w, http.StatusInternalServerError, err)
			return
		}
		data := modelsData.Data{
			Users: users,
			User:  user,
		}
		tmpl, err := template.ParseFiles("./ui/templates/all-users.html")
		if err != nil {
			pkg.ErrorLog.Println(err)
			pkg.Error(w, http.StatusInternalServerError, err)
			return
		}
		tmpl.Execute(w, data)
	} else if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			pkg.ErrorLog.Println(err)
			pkg.Error(w, http.StatusInternalServerError, err)
			return
		}
	} else {
		pkg.Error(w, http.StatusMethodNotAllowed, nil)
		return
	}
}
