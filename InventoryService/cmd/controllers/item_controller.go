package controllers

import (
	"inventoryService/m/v2/ent"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ItemController struct {
	Client *ent.Client
}

type Response struct {
	Msg  string      `json:"message"`
	Data interface{} `json:"data"`
}

func ResponseJSON(c *gin.Context, httpCode int, msg string, data interface{}) {
	c.JSON(httpCode, Response{
		Msg:  msg,
		Data: data,
	})
}

// @BasePath /api/v1

// GetIngridientsList godoc
// @Summary Gets list of all ingridients
// @Schemes
// @Description Get complete list of all ingridients availible for user
// @Tags example
// @Accept json
// @Produce application/json
// @Success 200
// @Router /inventory [get]
func (controller *ItemController) GetIngridientsList(c *gin.Context) {
	ingridients, err := controller.Client.Ingridient.Query().All(c)

	print("ingridient")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ResponseJSON(c, http.StatusOK, "", ingridients)
}

// @BasePath /api/v1

// AddIngridient godoc
// @Summary 	Adds new ingridient
// @Description Add new ingridient to database
// @Tags 		example
// @Accept 		json
// @Produce 	application/json
// @Param 		username body M true "username"
// @Success 	200 {string} Helloworld
// @Router  	/ingridient [post]
func (controller *ItemController) AddIngridient(c *gin.Context) {

	ingridient, err := controller.Client.Ingridient.
		Create().
		SetName("corona").
		SetAlcohol(0.05).
		SetMeasurmentUnit("ml").
		SetIsDry(false).
		Save(c)

	print(ingridient)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ResponseJSON(c, http.StatusOK, "", ingridient)
}
