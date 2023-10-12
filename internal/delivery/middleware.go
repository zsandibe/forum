package delivery

import (
	"context"
	"forum/pkg"
	"log"
	"net/http"

	modelsData "forum/internal/models/data"
	modelsUser "forum/internal/models/user"
)

// const contextKeyUser contentKey = "user"

// type contentKey string

func (h *Handler) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pkg.InfoLog.Printf("%s - %s %s %s", r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI())
		next.ServeHTTP(w, r)
	})
}

func (h *Handler) Middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session")
		var user modelsUser.User
		if err == http.ErrNoCookie {
			user = modelsUser.User{}
		} else if err == nil {
			user, err = h.service.Auth.UserByToken(cookie.Value)
			// fmt.Println("user: ", user)
			if err != nil {
				log.Printf("user by token: %s\n", err)
			}
		} else {
			pkg.Error(w, http.StatusBadRequest, nil)
		}
		// fmt.Println("DONE")
		// fmt.Println(user)
		ctx := context.WithValue(r.Context(), modelsData.ContextKeyUser, user)
		// fmt.Println(ctx)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
