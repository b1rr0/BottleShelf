package controllers

import (
	"context"
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

func NewUsersController(persister persistence.Persister) *UsersController {
	controller := new(UsersController)
	controller.persister = persister
	return controller
}

func (controller *UsersController) CreateUser(ctx context.Context, req models.CreateUserRequest) (resp models.CreateUserResponse, serr models.ServiceError) {
	res, err := controller.persister.SearchUsername(ctx, req.Username)
	if err != nil {
		return resp, models.NewInternalError(err)
	}
	if res {
		return resp, models.NewServiceError(http.StatusBadRequest, resources.UsernameTaken)
	}
	user := models.User{}
	user.Id = uuid.New()
	user.Username = req.Username
	user.Password = req.Password // TODO: Encrypt
	id, err := controller.persister.CreateUser(ctx, user)
	if err != nil {
		return resp, models.NewInternalError(err)
	}
	return models.CreateUserResponse{Id: id}, models.NoError()
}

func (controller *UsersController) CheckUser(ctx context.Context, req models.CheckUserRequest) (resp models.CheckUserResponse, serr models.ServiceError) {
	user := models.User{}
	user.Username = req.Username
	user.Password = req.Password // TODO: Encrypt
	id, err := controller.persister.CheckUser(ctx, user)
	if err != nil {
		return resp, models.NewInternalError(err)
	}
	return models.CheckUserResponse{Id: id}, models.NoError()
}

func (controller *UsersController) GetAllUsers(ctx context.Context) (resp models.GetUsersResponse, serr models.ServiceError) {
	users, err := controller.persister.GetAllUsers(ctx)
	if err != nil {
		return resp, models.NewInternalError(err)
	}
	usernames := []string{}
	for _, user := range users {
		usernames = append(usernames, user.Username)
	}
	return models.GetUsersResponse{Usernames: usernames}, models.NoError()
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
	resp, serr := controller.CreateUser(req.Context(), request)
	if !serr.IsOk() {
		serialization.SerializeServiceError(writer, serr)
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
	resp, serr := controller.CheckUser(req.Context(), request)
	if !serr.IsOk() {
		serialization.SerializeServiceError(writer, serr)
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
	resp, serr := controller.GetAllUsers(req.Context())
	if !serr.IsOk() {
		serialization.SerializeServiceError(writer, serr)
		return
	}
	// Prepare response
	serialization.SerializeResponse(writer, resp)
}
