package models

import (
	"net/http"
	"strings"

	"users_service/resources"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

type ServiceError struct {
	Status  int
	Message string
}

func (err *ServiceError) Error() string {
	return err.Message
}

func (err *ServiceError) Response() ErrorResponse {
	return ErrorResponse{Message: err.Message}
}

func (err *ServiceError) IsOk() bool {
	return err.Status == http.StatusOK || err.Status == http.StatusAccepted
}

func NewInternalError(err error) ServiceError {
	return ServiceError{Status: http.StatusInternalServerError, Message: strings.Join([]string{resources.InternalError, err.Error()}, ": ")}
}

func NewServiceError(status int, message string) ServiceError {
	return ServiceError{Status: status, Message: message}
}

func NoError() ServiceError {
	return ServiceError{Status: http.StatusOK, Message: ""}
}
