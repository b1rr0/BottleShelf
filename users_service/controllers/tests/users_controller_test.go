package tests

import (
	"context"
	"testing"
	"users_service/controllers"
	"users_service/models"
	"users_service/persistence"

	"github.com/google/uuid"
)

func TestNoUsers(t *testing.T) {
	p := persistence.NewLocalPersister()
	controller := controllers.NewUsersController(p)
	_, serr := controller.CheckUser(context.TODO(), models.CheckUserRequest{Username: "Shrek", Password: ""})
	if serr.IsOk() {
		t.Fatal("CheckUser succeded when no users exit")
	}
}

func TestCreateOneUser(t *testing.T) {
	controller := controllers.NewUsersController(persistence.NewLocalPersister())
	user1, serr := controller.CreateUser(context.TODO(), models.CreateUserRequest{Username: "Shrek", Password: "Kek"})
	if !serr.IsOk() {
		t.Fatal("Failed to create a user: ", serr)
	}
	if (user1.Id == uuid.UUID{}) {
		t.Fatal("User was created with invalid id")
	}
	user2, serr := controller.CheckUser(context.TODO(), models.CheckUserRequest{Username: "Shrek", Password: "Kek"})
	if !serr.IsOk() {
		t.Fatal("Faield to check user: ", serr)
	}
	if user1.Id != user2.Id {
		t.Fatal("User ids don't match")
	}
}

func TestGetUsers(t *testing.T) {
	controller := controllers.NewUsersController(persistence.NewLocalPersister())
	_, serr := controller.CreateUser(context.TODO(), models.CreateUserRequest{Username: "Shrek", Password: "Kek"})
	_, serr = controller.CreateUser(context.TODO(), models.CreateUserRequest{Username: "Fiona", Password: "Pass"})
	res, serr := controller.GetAllUsers(context.TODO())
	if !serr.IsOk() {
		t.Fatal("Failed to get all users: ", serr)
	}
	if len(res.Usernames) != 2 {
		t.Fatal("Returned invalid number of users")
	}
}
