package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// GetIngridients godoc
// @Summary Gets list of all ingridients
// @Schemes
// @Description Get complete list of all ingridients availible for user
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /inventory [get]
func GetIngridientsList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
