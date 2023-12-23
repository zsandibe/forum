package auth

import (
	"database/sql"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"

	s "forum/internal/models/session"
	modelsUser "forum/internal/models/user"
	u "forum/internal/models/user"
	p "forum/pkg"
)

func (a *AuthService) CreateUser(user u.User) error {
	// proverka na unique login
	if _, err := a.repo.GetUserByData(user.Username); err != sql.ErrNoRows {
		if err == nil {
			fmt.Println("ERROR username taken")
			return ErrUsernameTaken
		}
		fmt.Println("ERROR get user by data")
		return err
	}
	if user.AuthMethod == "" {
		if err := p.CheckUserInfo(user); err != nil {
			return err
		}

		password, err := p.GeneratePasswordHash(user.Password)
		if err != nil {
			return err
		}
		user.Password = password
	} else {
		if !p.CheckUsername(user.Username) {

			return p.ErrInvalidUsername
		}
	}

	if err := a.repo.CreateUser(user); err != nil {
		if strings.Contains(err.Error(), "UNIQUE") {
			return ErrEmailTaken
		}
		return err
	}
	return nil
}

func (a *AuthService) UserByToken(token string) (u.User, error) {
	// fmt.Println("OK")
	user, err := a.repo.UserByToken(token)
	// fmt.Println("user: ", user)
	if err != nil && err != sql.ErrNoRows {
		return user, nil
	}
	return user, nil
}

func (a *AuthService) CheckUser(u *u.User) (u.User, error) {
	if u.AuthMethod == "" {
		fmt.Println(u.Email)
		user, err := a.repo.GetUserByData(u.Email)
		fmt.Println(user.Email)
		if err != nil {
			return user, ErrNoUser
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password)); err != nil {
			return user, ErrWrongPassword
		}

		return user, nil
	}
	user, err := a.repo.GetUserByData(u.Email)
	fmt.Println(user, "user")
	if err != nil {
		return user, ErrNoUser
	}
	return user, nil
}

//--------------------------------------------------------------------------------------------

func (a *AuthService) generateToken() (string, error) {
	const tokenLength = 32
	bytes := make([]byte, tokenLength)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

//---------------------------------------------------------------------------------------------

func (a *AuthService) SetSession(user *u.User) (s.Session, error) {
	checkedUser, err := a.CheckUser(user)
	if err != nil {
		return s.Session{}, err
	}

	a.repo.DeleteSessionByUserID(checkedUser.ID)
	token, err := a.generateToken()
	// fmt.Println(token, "token: ")
	if err != nil {
		return s.Session{}, fmt.Errorf("failed to generate token: %w", err)
	}

	session := s.Session{
		UserID:      checkedUser.ID,
		Token:       token,
		ExpiresDate: time.Now().Add(sessionTime),
	}

	if err = a.repo.CreateSession(session); err != nil {
		// fmt.Println("OK")
		return session, fmt.Errorf("failed to create session: %w", err)
	}
	return session, nil
}

func (a *AuthService) DeleteSession(token string) error {
	return a.repo.DeleteSession(token)
}

func (a *AuthService) GetUserByID(userID int) (modelsUser.User, error) {
	return a.repo.GetUserByID(userID)
}

func (a *AuthService) GetAllUsersList() ([]u.User, error) {
	return a.repo.GetAllUsersList()
}
