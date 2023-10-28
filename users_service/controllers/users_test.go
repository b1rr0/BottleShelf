package controllers

import (
	"testing"
	"users_service/models"
	"users_service/persistence"
)

func TestNoUsers(t *testing.T) {
	p := persistence.NewLocalPersister()
	controller := NewUsersController(p)
	result := controller.CheckUser(models.CheckUser{Name: "Shrek", Password: ""})
	if result == true {
		t.Fatal("CheckUser returned true when no users exit")
	}
}

func TestCreateOneUser(t *testing.T) {
	controller := NewUsersController(persistence.NewLocalPersister())
	controller.CreateUser(models.CreateUser{Name: "Shrek", Password: "Kek"})
	result := controller.CheckUser(models.CheckUser{Name: "Shrek", Password: ""})
	if result == false {
		t.Fatal("CheckUser returned false for existing user")
	}
}
