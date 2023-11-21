package models

import (
	"github.com/google/uuid"
)

type CreateOrgRequest struct {
	OwnerId uuid.UUID `json:"owner"`
	Name    string    `json:"name"`
}

type CreateOrgResponse struct {
	Id uuid.UUID `json:"id"`
}

func (request *CreateOrgRequest) Validate() bool {
	return ValidateOrgname(request.Name)
}
