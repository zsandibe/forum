package auth

import (
	"database/sql"
	"fmt"

	s "forum/internal/models/session"
	u "forum/internal/models/user"
)

// var (
// 	user    u.User
// 	session s.Session
// )

func (r *AuthSql) CreateUser(user u.User) error {
	query := `
		INSERT INTO users (Username,Email,Password,Method,User_type,Requested) VALUES ($1, $2, $3, $4, $5, $6);
	`

	if _, err := r.db.Exec(query, user.Username, user.Email, user.Password, user.AuthMethod, user.Role, false); err != nil {
		return err
	}

	return nil
}

func (r *AuthSql) GetUserByData(userData string) (u.User, error) {
	var user u.User
	var method sql.NullString
	query := `
		SELECT ID,Username,Email,Password,Method,User_type FROM users WHERE  Email = ?;
	`
	row := r.db.QueryRow(query, userData)
	if err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &method, &user.Role); err != nil {
		fmt.Println(err)
		return u.User{}, err
	}
	user.AuthMethod = method.String
	fmt.Println(user, "user")
	return user, nil
}

func (r *AuthSql) CreateSession(sessions s.Session) error {
	query := `
        INSERT INTO sessions (UserID,Token,ExpDate) VALUES ($1, $2, $3);
    `

	if _, err := r.db.Exec(query, sessions.UserID, sessions.Token, sessions.ExpiresDate); err != nil {
		fmt.Println(err, "error creating session")
		return err
	}

	return nil
}

func (r *AuthSql) GetSession(token string) (s.Session, error) {
	query := `
		SELECT ID, UserID, Token, ExpiresDate FROM sessions WHERE Token = ?;
	`
	var session s.Session

	if err := r.db.QueryRow(query, token).Scan(&session.ID, &session.UserID, &session.Token, &session.ExpiresDate); err != nil {
		return session, err
	}

	return session, nil
}

func (r *AuthSql) DeleteSession(token string) error {
	query := `
        DELETE FROM sessions WHERE Token = ?;
    `

	if _, err := r.db.Exec(query, token); err != nil {
		return err
	}

	return nil
}

func (r *AuthSql) DeleteSessionByUserID(userID int) error {
	query := `
		DELETE FROM sessions WHERE UserID = ?;
	`

	if _, err := r.db.Exec(query, userID); err != nil {
		return err
	}
	return nil
}

func (r *AuthSql) UserByToken(token string) (u.User, error) {
	query := `
		SELECT users.ID, users.Username, users.Email, users.Password ,users.User_type
		FROM sessions INNER JOIN users
		ON users.ID = sessions.UserID
		WHERE sessions.Token = ?;
	`
	var user u.User

	if err := r.db.QueryRow(query, token).Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Role); err != nil {
		return user, err
	}
	return user, nil
}

func (r *AuthSql) GetUserByID(userID int) (u.User, error) {
	var user u.User
	query := `
		SELECT ID, Username, Email,Password FROM users WHERE ID = ?;
	`

	if err := r.db.QueryRow(query, userID).Scan(&user.ID, &user.Username, &user.Email, &user.Password); err != nil {
		return u.User{}, err
	}
	return user, nil
}

func (r *AuthSql) GetAllUsersList() ([]u.User, error) {
	var users []u.User
	query := `
		SELECT users.ID,users.Username,users.Email,users.User_type,users.Requested FROM users
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return users, err
	}
	defer rows.Close()

	for rows.Next() {
		var user u.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Role, &user.Requested); err != nil {
			return nil, fmt.Errorf("can`t get users : %w", err)
		}
		users = append(users, user)
	}
	return users, nil
}
