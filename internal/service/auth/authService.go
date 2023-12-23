package auth

import (
	modelsSession "forum/internal/models/session"
	modelsUser "forum/internal/models/user"
	u "forum/internal/models/user"
	repositoryAuth "forum/internal/repository/auth"
)

type Auth interface {
	CreateUser(user modelsUser.User) error
	GetUserByID(userID int) (modelsUser.User, error)
	SetSession(user *modelsUser.User) (modelsSession.Session, error)
	DeleteSession(token string) error
	UserByToken(token string) (modelsUser.User, error)
	GetAllUsersList() ([]u.User, error)
}

type AuthService struct {
	repo repositoryAuth.Authorization
}

func NewAuthService(repo repositoryAuth.Authorization) *AuthService {
	return &AuthService{
		repo: repo,
	}
}
