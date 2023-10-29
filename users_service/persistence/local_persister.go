package persistence

import (
	"users_service/models"
)

type LocalPersister struct {
	users []models.User
}

func (persister *LocalPersister) CreateUser(user models.User) bool {
	persister.users = append(persister.users, user)
	return true
}

func (persister *LocalPersister) SearchUsername(name string) bool {
	for _, user := range persister.users {
		if user.Username == name {
			return true
		}
	}
	return false
}

func (persister *LocalPersister) CheckUser(checkUser *models.User) bool {
	for _, user := range persister.users {
		if user.Username == checkUser.Username && user.Password == checkUser.Password {
			*checkUser = user
			return true
		}
	}
	return false
}

func NewLocalPersister() *LocalPersister {
	return new(LocalPersister)
}
