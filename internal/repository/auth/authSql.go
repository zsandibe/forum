package auth

import (
	"database/sql"
	s "forum/internal/models/session"
	u "forum/internal/models/user"
)

type AuthSql struct {
	db *sql.DB
}

type Authorization interface {
	CreateUser(user u.User) error
	GetUserByData(userData string) (u.User, error)
	GetUserByID(userID int) (u.User, error)
	GetSession(token string) (s.Session, error)
	CreateSession(sessions s.Session) error
	DeleteSession(token string) error
	DeleteSessionByUserID(userID int) error
	UserByToken(token string) (u.User, error)
}

func NewAuthSql(db *sql.DB) *AuthSql {
	return &AuthSql{
		db: db,
	}
}
