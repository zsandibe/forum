package pkg

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func GeneratePasswordHash(password string) (string, error) {
	pass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	fmt.Println(string(pass))
	return string(pass), err
}
