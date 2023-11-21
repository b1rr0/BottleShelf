package models

import (
	"github.com/google/uuid"
)

type Org struct {
	Id   uuid.UUID
	Name string
}
