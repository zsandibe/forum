package request

import (
	requestModels "forum/internal/models/request"
)

type requestProvider interface {
	CreateRequest(request requestModels.Request) (int64, error)
	GetRequestByUserID(userID int) (requestModels.Request, error)
	GetAllRequests() ([]requestModels.Request, error)
	UpdateRequestStatus(requestID, requestUserID int, status string) error
	UpdateUserRole(userId int, role string) error
}

type RequestHandler struct {
	request requestProvider
}

func NewRequestHandler(request requestProvider) *RequestHandler {
	return &RequestHandler{
		request: request,
	}
}
