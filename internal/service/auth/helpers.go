package auth

import (
	"errors"
	"time"
)

var (
	ErrNoUser        = errors.New("user not found")
	ErrWrongPassword = errors.New("wrong password")
	ErrUsernameTaken = errors.New("username already taken")
	ErrEmailTaken    = errors.New("email already taken")
)

const sessionTime = time.Hour * 6
