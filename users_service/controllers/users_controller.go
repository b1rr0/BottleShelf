package controllers

import (
	"errors"
	"users_service/models"
	"users_service/persistence"

	"github.com/google/uuid"
)

type UsersController struct {
	persister persistence.UsersPersister
}

func (controller *UsersController) CreateUser(req models.CreateUserRequest) (resp models.CreateUserResponse, err error) {
	if controller.persister.SearchUsername(req.Username) {
		err = errors.New("This username is already taken")
		return
	}
	user := models.User{}
	user.Id = uuid.New()
	user.Username = req.Username
	user.Password = req.Password // TODO: Encrypt
	if !controller.persister.CreateUser(user) {
		err = errors.New("Failed to create a user")
		return
	}
	resp = models.CreateUserResponse{Id: user.Id}
	err = nil
	return
}

func (controller *UsersController) CheckUser(req models.CheckUserRequest) (resp models.CheckUserResponse, err error) {
	user := models.User{}
	user.Username = req.Username
	user.Password = req.Password // TODO: Encrypt
	if !controller.persister.CheckUser(&user) {
		err = errors.New("User not found")
		return
	}
	resp = models.CheckUserResponse{Id: user.Id}
	err = nil
	return
}

func NewUsersController(persister persistence.UsersPersister) *UsersController {
	controller := new(UsersController)
	controller.persister = persister
	return controller
}
