package controllers

import (
	"users_service/models"
	"users_service/persistence"
)

type UsersController struct {
	persister persistence.UsersPersister
}

func (controller *UsersController) CreateUser(req models.CreateUser) {
	controller.persister.CreateUser(models.User{Name: req.Name, Password: req.Password})
}

func (controller *UsersController) CheckUser(req models.CheckUser) bool {
	return controller.persister.CheckUser(req.Name)
}

func NewUsersController(persister persistence.UsersPersister) *UsersController {
	controller := new(UsersController)
	controller.persister = persister
	return controller
}
