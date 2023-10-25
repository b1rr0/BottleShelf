package persistence

import (
	"usersService/m/v2/models"
)

type LocalPersister struct {
	users []models.User
}

func (persister *LocalPersister) CreateUser(user models.User) {
	persister.users = append(persister.users, user)
}

func (persister *LocalPersister) CheckUser(name string) bool {
	for _, user := range persister.users {
		if user.Name == name {
			return true
		}
	}
	return false
}

func NewLocalPersister() *LocalPersister {
	return new(LocalPersister)
}
