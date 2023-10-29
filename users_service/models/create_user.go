package models

import (
	"github.com/google/uuid"
)

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateUserResponse struct {
	Id uuid.UUID `json:"id"`
}

func (request *CreateUserRequest) Validate() bool {
	return ValidateUsername(request.Username) && ValidatePassword(request.Password)
}
