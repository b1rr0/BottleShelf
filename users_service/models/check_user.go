package models

import (
	"github.com/google/uuid"
)

type CheckUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CheckUserResponse struct {
	Id uuid.UUID `json:"id"`
}

func (request *CheckUserRequest) Validate() bool {
	return ValidateUsername(request.Username)
}
