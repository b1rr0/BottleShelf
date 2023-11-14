package controllers

import (
	"context"
	"errors"
	"net/http"

	"users_service/models"
	"users_service/persistence"
	"users_service/resources"
	"users_service/serialization"

	"github.com/google/uuid"
)

type UsersController struct {
	persister persistence.Persister
}

func (controller *UsersController) CreateUser(ctx context.Context, req models.CreateUserRequest) (resp models.CreateUserResponse, err error) {
	res, err := controller.persister.SearchUsername(ctx, req.Username)
	if err != nil {
		return
	}
	if res {
		err = errors.New("This username is already taken")
		return
	}
	user := models.User{}
	user.Id = uuid.New()
	user.Username = req.Username
	user.Password = req.Password // TODO: Encrypt
	id, err := controller.persister.CreateUser(ctx, user)
	if err != nil {
		return
	}
	resp = models.CreateUserResponse{Id: id}
	err = nil
	return
}

func (controller *UsersController) CheckUser(ctx context.Context, req models.CheckUserRequest) (resp models.CheckUserResponse, err error) {
	user := models.User{}
	user.Username = req.Username
	user.Password = req.Password // TODO: Encrypt
	id, err := controller.persister.CheckUser(ctx, user)
	if err != nil {
		// TODO: Return right error string depending on error
		return
	}
	resp = models.CheckUserResponse{Id: id}
	err = nil
	return
}

func (controller *UsersController) GetAllUsers(ctx context.Context) (resp models.GetUsersResponse, err error) {
	users, err := controller.persister.GetAllUsers(ctx)
	if err != nil {
		return
	}
	usernames := []string{}
	for _, user := range users {
		usernames = append(usernames, user.Username)
	}
	resp = models.GetUsersResponse{Usernames: usernames}
	err = nil
	return
}

func NewUsersController(persister persistence.Persister) *UsersController {
	controller := new(UsersController)
	controller.persister = persister
	return controller
}

/*
API ENDPOINTS
*/

//  @Summary      Create a new user
//  @Accept       json
//  @Produce      json
//  @Param        user    body        models.CreateUserRequest    true    "Create user"
//  @Failure      400     {object}    models.ErrorResponse
//  @Success      200     {object}    models.CheckUserResponse
//  @Router       /users/create [post]
func (controller *UsersController) ServeCreateUser(writer http.ResponseWriter, req *http.Request) {
	// Parse request
	request := models.CreateUserRequest{}
	status := serialization.DeserializeRequest(req, &request)
	if status != http.StatusOK {
		serialization.SerializeError(writer, status, resources.SerializationFailed)
		return
	}
	// Call controller
	resp, err := controller.CreateUser(req.Context(), request)
	if err != nil {
		serialization.SerializeError(writer, http.StatusBadRequest, err.Error())
		return
	}
	// Prepare response
	serialization.SerializeResponse(writer, resp)
}

//  @Summary      Check user data (Log in)
//  @Accept       json
//  @Produce      json
//  @Param        user    body        models.CheckUserRequest    true    "Log in"
//  @Failure      400     {object}    models.ErrorResponse
//  @Failure      404     {object}    models.ErrorResponse
//  @Success      200     {object}    models.CheckUserResponse
//  @Router       /users/check [post]
func (controller *UsersController) ServeCheckUser(writer http.ResponseWriter, req *http.Request) {
	// Parse request
	request := models.CheckUserRequest{}
	status := serialization.DeserializeRequest(req, &request)
	if status != http.StatusOK {
		serialization.SerializeError(writer, status, resources.SerializationFailed)
		return
	}
	// Call controller
	resp, err := controller.CheckUser(req.Context(), request)
	if err != nil {
		serialization.SerializeError(writer, http.StatusNotFound, resources.WrongUserData)
		return
	}
	// Prepare response
	serialization.SerializeResponse(writer, resp)
}

//  @Summary      Get all usernames
//  @Description  Debug-only endpoint, should not be called from api gateway
//  @Produce      json
//  @Failure      400     {object}    models.ErrorResponse
//  @Success      200     {object}    models.GetUsersResponse
//  @Router       /users [get]
func (controller *UsersController) ServeAllUsernames(writer http.ResponseWriter, req *http.Request) {
	// Call controller
	resp, err := controller.GetAllUsers(req.Context())
	if err != nil {
		serialization.SerializeError(writer, http.StatusBadRequest, err.Error())
		return
	}
	// Prepare response
	serialization.SerializeResponse(writer, resp)
}
