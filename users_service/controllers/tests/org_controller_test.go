package tests

import (
	"context"
	"testing"
	"users_service/controllers"
	"users_service/models"
	"users_service/persistence"

	"github.com/google/uuid"
)

func TestCreateOneOrg(t *testing.T) {
	p := persistence.NewLocalPersister()
	orgController := controllers.NewOrgController(p)
	usersController := controllers.NewUsersController(p)

	user1, serr := usersController.CreateUser(context.TODO(), models.CreateUserRequest{Username: "Shrek", Password: "Kek"})
	org, serr := orgController.CreateOrg(context.TODO(), models.CreateOrgRequest{OwnerId: user1.Id, Name: "test"})
	if !serr.IsOk() {
		t.Fatal("Failed to create an org: ", serr)
	}
	if (org.Id == uuid.UUID{}) {
		t.Fatal("Org was created with invalid id")
	}
	orgs, serr := orgController.GetAllOrgs(context.TODO())
	if len(orgs.Orgnames) != 1 {
		t.Fatal("Wrong number of orgs returned")
	}
}

func TestCreateOrgInvalidOwner(t *testing.T) {
	orgController := controllers.NewOrgController(persistence.NewLocalPersister())
	_, serr := orgController.CreateOrg(context.TODO(), models.CreateOrgRequest{OwnerId: uuid.UUID{}, Name: "test"})
	if serr.IsOk() {
		t.Fatal("Created org with invalid owner")
	}
}
