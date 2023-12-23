package request

import (
	"database/sql"
	"errors"
	"fmt"
	requestModels "forum/internal/models/request"
)

func (s *RequestService) CreateRequest(request requestModels.Request) (int64, error) {
	if err := s.repo.DeleteRequestByUserId(request.UserID); err != nil {
		return 0, err
	}
	return s.repo.CreateRequest(request)
}

func (s *RequestService) GetRequestByUserID(userID int) (requestModels.Request, error) {
	request, err := s.repo.GetRequestByUserID(userID)
	if errors.Is(err, sql.ErrNoRows) {
		return requestModels.Request{}, errors.New("request not found")
	} else if err != nil {
		return requestModels.Request{}, err
	}
	return request, nil
}

func (s *RequestService) GetAllRequests() ([]requestModels.Request, error) {
	return s.repo.GetAllRequests()
}

func (s *RequestService) UpdateRequestStatus(requestID, requestUserID int, status string) error {
	if status == "accept" {
		err := s.repo.UpdateUserRole(requestUserID, "moderator")
		if err != nil {
			return fmt.Errorf("DB can't update user type: %w", err)
		}
	}
	return s.repo.UpdateRequestStatus(requestID, status)
}

func (s *RequestService) UpdateUserRole(userId int, role string) error {
	return s.repo.UpdateUserRole(userId, role)
}

func (s *RequestService) DeleteRequestByUserId(userId int) error {
	return s.repo.DeleteRequestByUserId(userId)
}
