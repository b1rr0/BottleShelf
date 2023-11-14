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
	_, err := controller.CheckUser(context.TODO(), models.CheckUserRequest{Username: "Shrek", Password: ""})
	if err == nil {
		t.Fatal("CheckUser succeded when no users exit")
	}
}

func TestCreateOneUser(t *testing.T) {
	controller := controllers.NewUsersController(persistence.NewLocalPersister())
	user1, err := controller.CreateUser(context.TODO(), models.CreateUserRequest{Username: "Shrek", Password: "Kek"})
	if err != nil {
		t.Fatal("Failed to create a user: ", err)
	}
	if (user1.Id == uuid.UUID{}) {
		t.Fatal("User was created with invalid id")
	}
	user2, err := controller.CheckUser(context.TODO(), models.CheckUserRequest{Username: "Shrek", Password: "Kek"})
	if err != nil {
		t.Fatal("Faield to check user: ", err)
	}
	if user1.Id != user2.Id {
		t.Fatal("User ids don't match")
	}
}

func TestGetUsers(t *testing.T) {
	controller := controllers.NewUsersController(persistence.NewLocalPersister())
	_, err := controller.CreateUser(context.TODO(), models.CreateUserRequest{Username: "Shrek", Password: "Kek"})
	_, err = controller.CreateUser(context.TODO(), models.CreateUserRequest{Username: "Fiona", Password: "Pass"})
	res, err := controller.GetAllUsers(context.TODO())
	if err != nil {
		t.Fatal("Failed to get all users: ", err)
	}
	if len(res.Usernames) != 2 {
		t.Fatal("Returned invalid number of users")
	}
}
