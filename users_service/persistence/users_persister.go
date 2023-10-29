package persistence

import (
	"users_service/models"
)

type UsersPersister interface {
	CreateUser(user models.User) bool
	CheckUser(user *models.User) bool
	SearchUsername(name string) bool
}
