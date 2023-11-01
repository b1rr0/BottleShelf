package tests

import (
	"testing"
	"users_service/controllers"
	"users_service/models"
	"users_service/persistence"
)

func TestNoUsers(t *testing.T) {
	p := persistence.NewLocalPersister()
	controller := controllers.NewUsersController(p)
	_, err := controller.CheckUser(models.CheckUserRequest{Username: "Shrek", Password: ""})
	if err == nil {
		t.Fatal("CheckUser succeded when no users exit")
	}
}

func TestCreateOneUser(t *testing.T) {
	controller := controllers.NewUsersController(persistence.NewLocalPersister())
	controller.CreateUser(models.CreateUserRequest{Username: "Shrek", Password: "Kek"})
	_, err := controller.CheckUser(models.CheckUserRequest{Username: "Shrek", Password: "Kek"})
	if err != nil {
		t.Fatal("CheckUser returned false for existing user")
	}
}
