package session

import "time"

type Session struct {
	ID          int
	UserID      int
	Token       string
	ExpiresDate time.Time
}
