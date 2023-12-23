package request

import (
	"database/sql"
	requestModels "forum/internal/models/request"
)

type Request interface {
	CreateRequest(request requestModels.Request) (int64, error)
	GetRequestByUserID(UserID int) (requestModels.Request, error)
	GetAllRequests() ([]requestModels.Request, error)
	UpdateRequestStatus(requestID int, status string) error
	UpdateUserRole(userId int, role string) error
	DeleteRequestByUserId(userId int) error
}

type RequestSql struct {
	db *sql.DB
}

func NewRequestSql(db *sql.DB) *RequestSql {
	return &RequestSql{
		db: db,
	}
}
