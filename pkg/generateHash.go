package pkg

import "golang.org/x/crypto/bcrypt"

func GeneratePasswordHash(password string) (string, error) {
	pass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return string(pass), err
}
