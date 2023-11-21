package persistence

import (
	"context"
	"fmt"

	"users_service/models"

	"github.com/google/uuid"
)

type LocalPersister struct {
	users []models.User
	orgs  []models.Org
}

func NewLocalPersister() *LocalPersister {
	return new(LocalPersister)
}

func (persister *LocalPersister) CreateUser(ctx context.Context, user models.User) (uuid.UUID, error) {
	persister.users = append(persister.users, user)
	return user.Id, nil
}

func (persister *LocalPersister) SearchUsername(ctx context.Context, username string) (bool, error) {
	for _, user := range persister.users {
		if user.Username == username {
			return true, nil
		}
	}
	return false, nil
}

func (persister *LocalPersister) SearchUserId(ctx context.Context, id uuid.UUID) (bool, error) {
	for _, user := range persister.users {
		if user.Id == id {
			return true, nil
		}
	}
	return false, nil
}

func (persister *LocalPersister) CheckUser(ctx context.Context, checkUser models.User) (uuid.UUID, error) {
	for _, user := range persister.users {
		if user.Username == checkUser.Username && user.Password == checkUser.Password {
			return user.Id, nil
		}
	}
	return uuid.UUID{}, fmt.Errorf("Not found")
}

func (persister *LocalPersister) GetAllUsers(ctx context.Context) (users []models.User, err error) {
	for _, user := range persister.users {
		users = append(users, user)
	}
	return users, nil
}

func (persister *LocalPersister) CreateOrg(ctx context.Context, org models.Org, owner uuid.UUID) (uuid.UUID, error) {
	persister.orgs = append(persister.orgs, org)
	return org.Id, nil
}

func (persister *LocalPersister) GetAllOrgs(ctx context.Context) (orgs []models.Org, err error) {
	for _, org := range persister.orgs {
		orgs = append(orgs, org)
	}
	return orgs, nil
}
