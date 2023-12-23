package request

import (
	"fmt"
	requestModels "forum/internal/models/request"
)

func (r *RequestSql) CreateRequest(request requestModels.Request) (int64, error) {
	fmt.Println(request)
	query := `INSERT INTO requests 
	(UserID, Created_at, Status) values (?,?,?)`
	res, err := r.db.Exec(query, &request.UserID, &request.CreatedAt, &request.Status)
	if err != nil {
		return 0, err
	}
	id, _ := res.LastInsertId()
	return id, nil
}

func (r *RequestSql) GetRequestByUserID(userID int) (requestModels.Request, error) {
	var request requestModels.Request
	query := `
		SELECT requests.ID,requests.UserID, requests.Created_at, requests.Status FROM requests
		INNER JOIN users ON users.ID = requests.UserID
		WHERE users.ID = $1 
	`
	if err := r.db.QueryRow(query, userID).Scan(&request.ID, &request.UserID, &request.CreatedAt, &request.Status); err != nil {
		return request, err
	}

	return request, nil
}

func (r *RequestSql) GetAllRequests() ([]requestModels.Request, error) {
	var requests []requestModels.Request

	query := `
		SELECT requests.ID,requests.UserID, requests.Created_at, requests.Status FROM requests 
		INNER JOIN users ON users.ID = requests.UserID
		ORDER BY requests.ID DESC 
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var request requestModels.Request
		if err := rows.Scan(&request.ID, &request.UserID, &request.CreatedAt, &request.Status); err != nil {
			return requests, err
		}
		requests = append(requests, request)
	}
	return requests, nil
}

func (r *RequestSql) UpdateRequestStatus(requestID int, status string) error {
	query := `
		UPDATE requests SET status = ? WHERE ID = ?
	`
	if _, err := r.db.Exec(query, status, requestID); err != nil {
		return fmt.Errorf("can`t update request status: %w", err)
	}
	return nil
}

func (r *RequestSql) UpdateUserRole(userId int, role string) error {
	query := `UPDATE users SET User_type = ? WHERE ID = ?`

	if _, err := r.db.Exec(query, role, userId); err != nil {
		return fmt.Errorf("can't update user type: %w", err)
	}

	return nil
}

func (r *RequestSql) DeleteRequestByUserId(userId int) error {
	query := `
		DELETE FROM requests WHERE UserID = ?
	`
	if _, err := r.db.Exec(query, userId); err != nil {
		return err
	}
	return nil
}
