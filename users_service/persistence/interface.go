package persistence

import (
	"users_service/models"
)

type UsersPersister interface {
	CreateUser(user models.User)
	CheckUser(name string) bool
}
