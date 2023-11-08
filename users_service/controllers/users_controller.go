package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	"users_service/models"
	"users_service/persistence"
	"users_service/serialization"

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

//  @Summary      Create a new user
//  @Accept       json
//  @Produce      json
//  @Param        user    body    models.CreateUserRequest    true    "Create user"
//  @Failure      400     body    nil                                 "Bad request"
//  @Success      200
//  @Router       /users/create [post]
func (controller *UsersController) ServeCreateUser(writer http.ResponseWriter, req *http.Request) {
	// Parse request
	request := models.CreateUserRequest{}
	status := serialization.DeserializeRequest(req, &request)
	if status != http.StatusOK {
		writer.WriteHeader(status)
		return
	}
	// Call controller
	resp, err := controller.CreateUser(request)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte(err.Error()))
		return
	}
	// Prepare response
	marshalled, err := json.Marshal(resp)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Write(marshalled)
}

//  @Summary      Check user data (Log in)
//  @Accept       json
//  @Produce      json
//  @Param        user    body    models.CheckUserRequest    true    "Log in"
//  @Failure      400     body    nil                                "Bad request"
//  @Failure      404     body    nil                                "User not found"
//  @Success      200
//  @Router       /users/check [post]
func (controller *UsersController) ServeCheckUser(writer http.ResponseWriter, req *http.Request) {
	// Parse request
	request := models.CheckUserRequest{}
	status := serialization.DeserializeRequest(req, &request)
	if status != http.StatusOK {
		writer.WriteHeader(status)
		return
	}
	// Call controller
	resp, err := controller.CheckUser(request)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte(err.Error()))
		return
	}
	// Prepare response
	marshalled, err := json.Marshal(resp)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Write(marshalled)
}
