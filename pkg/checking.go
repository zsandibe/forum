package pkg

import (
	"errors"
	"fmt"
	u "forum/internal/models/user"
	"regexp"
	"strings"
)

var (
	ErrInvalidEmail    = errors.New("invalid email address")
	ErrInvalidPassword = errors.New("invalid password")
	ErrInvalidUsername = errors.New("invalid username")
)

func CheckUserInfo(user u.User) error {
	// proverka email

	if !regexp.MustCompile(`^([a-zA-Z0-9._%-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4})$`).MatchString(user.Email) {
		return ErrInvalidEmail
	}

	if !CheckUsername(user.Username) {
		return ErrInvalidUsername
	}

	if !CheckPassword(user.Password) {
		return ErrInvalidPassword
	}
	return nil
}

func CheckUsername(username string) bool {
	for _, val := range username {
		if val < 32 || val > 126 {
			return false
		}
	}
	return true
}

func CheckPassword(password string) bool {
	nums := "0123456789"
	lowCase := "abcdefghijklmnopqrstuvwxyz"
	upCase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// symbols := "!@#$%^&*()_-+={[}]|\\:;<,>.?/"

	if !(len(password) > 8 && len(password) < 20) {
		return false
	}

	if !contains(password, nums) || !contains(password, lowCase) || !contains(password, upCase)  {
		fmt.Println("Password contains")
		return false
	}

	// for _, val := range symbols {
	// 	if val < 32 || val > 126 {
	// 		return false
	// 	}
	// }
	return true
}

func contains(password, numbers string) bool {
	for _, val := range numbers {
		if strings.Contains(password, string(val)) {
			return true
		}
	}
	return false
}
