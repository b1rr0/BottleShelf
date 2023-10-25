package persistence

import (
	"usersService/m/v2/models"
)

type UsersPersister interface {
	CreateUser(user models.User)
	CheckUser(name string) bool
}
