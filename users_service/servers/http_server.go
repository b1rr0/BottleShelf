package servers

import (
	"encoding/json"
	"net/http"

	"users_service/controllers"
	"users_service/models"
)

type HttpServer struct {
	usersController *controllers.UsersController
}

func NewHttpServer(usersController *controllers.UsersController) *HttpServer {
	server := new(HttpServer)
	server.usersController = usersController
	return server
}

//  @Summary      Get information if service is healthy
//  @Failure      default
//  @Success      200
//  @Router       /health [get]
func (server *HttpServer) ServeHealth(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Content-Type", "text/plain")
	writer.Write([]byte("Alive\n"))
}

//  @Summary      Create a new user
//  @Accept       json
//  @Produce      json
//  @Param		  user	        body		models.CreateUserRequest	    true     "Create user"
//  @Failure      400
//  @Success      200
//  @Router       /users/create [post]
func (server *HttpServer) ServeCreateUser(writer http.ResponseWriter, req *http.Request) {
	// Parse request
	request := models.CreateUserRequest{}
	status := DeserializeRequest(req, &request)
	if status != http.StatusOK {
		writer.WriteHeader(status)
		return
	}
	// Call controller
	resp, err := server.usersController.CreateUser(request)
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
//  @Param		  user	        body		models.CheckUserRequest	    true     "Log in"
//  @Failure      404
//  @Success      200
//  @Router       /users/check [post]
func (server *HttpServer) ServeCheckUser(writer http.ResponseWriter, req *http.Request) {
	// Parse request
	request := models.CheckUserRequest{}
	status := DeserializeRequest(req, &request)
	if status != http.StatusOK {
		writer.WriteHeader(status)
		return
	}
	// Call controller
	resp, err := server.usersController.CheckUser(request)
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
