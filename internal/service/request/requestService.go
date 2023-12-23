package request

import (
	requestModels "forum/internal/models/request"
	repositoryRequest "forum/internal/repository/request"
)

type Request interface {
	CreateRequest(request requestModels.Request) (int64, error)
	GetRequestByUserID(userID int) (requestModels.Request, error)
	GetAllRequests() ([]requestModels.Request, error)
	UpdateRequestStatus(requestID, requestUserID int, status string) error
	UpdateUserRole(userId int, role string) error
	DeleteRequestByUserId(userId int) error
}

type RequestService struct {
	repo repositoryRequest.Request
}

func NewRequestService(repo repositoryRequest.Request) *RequestService {
	return &RequestService{
		repo: repo,
	}
}
