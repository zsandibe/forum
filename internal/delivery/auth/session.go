package auth

import (
	u "forum/internal/models/user"
	"net/http"
)

func (a *AuthHandler) SetSession(w http.ResponseWriter, user *u.User) error {
	session, err := a.user.SetSession(user)
	// fmt.Println(err)
	if err != nil {
		return err
	}
	cookie := &http.Cookie{
		Name:     "session",
		Value:    session.Token,
		Path:     "/",
		Expires:  session.ExpiresDate,
		HttpOnly: true,
		MaxAge:   3600,
	}
	http.SetCookie(w, cookie)
	return nil
}
