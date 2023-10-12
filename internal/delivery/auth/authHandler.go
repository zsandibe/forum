package auth

import (
	s "forum/internal/models/session"
	u "forum/internal/models/user"
)

type AuthProvider interface {
	CreateUser(user u.User) error
	SetSession(user *u.User) (s.Session, error)
	DeleteSession(token string) error
}

type AuthHandler struct {
	user AuthProvider
}

func NewAuthHandler(user AuthProvider) *AuthHandler {
	return &AuthHandler{
		user: user,
	}
}
