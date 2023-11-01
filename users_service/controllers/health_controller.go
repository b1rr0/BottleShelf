package controllers

import (
	"net/http"
)

type HealthController struct {
}

//  @Summary      Get information if service is healthy
//  @Failure      default
//  @Success      200
//  @Router       /health [get]
func (controller *HealthController) ServeHealth(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Content-Type", "text/plain")
	writer.Write([]byte("Alive\n"))
}
