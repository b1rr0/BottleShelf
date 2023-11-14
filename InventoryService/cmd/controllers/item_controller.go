package controllers

import (
	"encoding/json"
	"inventoryService/m/v2/cmd/models"
	"inventoryService/m/v2/cmd/resources"
	"inventoryService/m/v2/ent"
	"inventoryService/m/v2/ent/ingridient"
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

func (controller *ItemController) ValidateNewIngridient(c *gin.Context, item models.ItemModelCreate) (passed bool, errorMessage string) {
	isAlreadyExists, err := controller.Client.Ingridient.
		Query().
		Where(ingridient.Name(item.Name)).
		Exist(c)

	errorMessage = ""

	if err != nil {
		ResponseJSON(c, http.StatusInternalServerError, err.Error(), nil)
		passed = false
		errorMessage += err.Error()
		return
	}

	if isAlreadyExists {
		errorMessage += resources.AlreadyExists + " "
		passed = false
	}

	if item.Alcohol < 0 || item.Alcohol > 1 {
		errorMessage += resources.InadequateAlcohol + " "
		passed = false
	}

	if !(item.MeasurmentUnit == "pcs" || item.MeasurmentUnit == "g" || item.MeasurmentUnit == "ml" || item.MeasurmentUnit == "") {
		errorMessage += resources.WrongMeasurement + " "
		passed = false
	}
	return
}

// @BasePath /api/v1

// GetIngridientsList godoc
// @Summary        Gets list of all ingridients
// @Description    Get complete list of all ingridients availible for user
// @Tags           Inventory manipulation
// @Produce        application/json
// @Success        200
// @Router         /inventory [get]
func (controller *ItemController) GetIngridientsList(c *gin.Context) {
	ingridients, err := controller.Client.Ingridient.
		Query().
		All(c)

	if err != nil {
		ResponseJSON(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	ResponseJSON(c, http.StatusOK, "", ingridients)
}

// @BasePath /api/v1

// GetIngridientsList godoc
// @Summary         Gets list ingridients by filter
// @Description     Get list of ingridients filtering by it's name and/or parameters
// @Tags            Inventory manipulation
// @Produce         application/json
// @Param           item query models.ItemModelFilters true "Item to search for"
// @Success         200
// @Router          /ingridient/search [get]
func (controller *ItemController) GetIngridientByFilter(c *gin.Context) {
	var filters models.ItemModelFilters
	filters.AlcoholMax = 1

	err := c.BindQuery(&filters)

	if err != nil {
		ResponseJSON(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if filters.AlcoholMin < 0 || filters.AlcoholMax > 1 {
		ResponseJSON(c, http.StatusBadRequest, resources.InadequateAlcohol, nil)
		return
	}

	ingridients, err := controller.Client.Ingridient.
		Query().
		Where(
			ingridient.And(
				ingridient.NameContainsFold(filters.Name),
				ingridient.AlcoholGT(filters.AlcoholMin),
				ingridient.AlcoholLT(filters.AlcoholMax),
				ingridient.IsDry(filters.IsDry),
			),
		).
		All(c)

	if err != nil {
		ResponseJSON(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	ResponseJSON(c, http.StatusOK, "", ingridients)
}

// @BasePath /api/v1

// AddIngridient godoc
// @Summary           Adds new ingridient
// @Description       Add new ingridient to database
// @Tags              Inventory manipulation
// @Accept            json
// @Produce           application/json
// @Param             item body models.ItemModelCreate true "Item data"
// @Success           201
// @Router            /ingridient [post]
func (controller *ItemController) AddIngridient(c *gin.Context) {
	jsonData, err := c.GetRawData()
	if err != nil {
		ResponseJSON(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	var item models.ItemModelCreate
	json.Unmarshal(jsonData, &item)

	passedValidation, errorMessage := controller.ValidateNewIngridient(c, item)

	if !passedValidation {
		ResponseJSON(c, http.StatusBadRequest, errorMessage, nil)
		return
	}

	ingridient, err := controller.Client.Ingridient.
		Create().
		SetName(item.Name).
		SetAlcohol(item.Alcohol).
		SetMeasurmentUnit(ingridient.MeasurmentUnit(item.MeasurmentUnit)).
		SetIsDry(item.IsDry).
		Save(c)

	if err != nil {
		ResponseJSON(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	ResponseJSON(c, http.StatusCreated, resources.IngridientAdded, ingridient)
}

// @BasePath /api/v1

// AddIngridient godoc
// @Summary           Changes ingridient information
// @Description       Change ingridient in the database by id
// @Tags              Inventory manipulation
// @Accept            json
// @Produce           application/json
// @Param             item body models.ItemModel true "Item id and it's new data"
// @Success           202
// @Router            /ingridient [put]
func (controller *ItemController) ChangeIngridient(c *gin.Context) {
	jsonData, err := c.GetRawData()
	if err != nil {
		ResponseJSON(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	var item models.ItemModel
	json.Unmarshal(jsonData, &item)

	ingridientOld, err := controller.Client.Ingridient.
		Query().
		Where(ingridient.ID(item.Id)).
		Only(c)

	if err != nil {
		if ingridientOld == nil {
			ResponseJSON(c, http.StatusNotFound, err.Error(), nil)
			return
		}
		ResponseJSON(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	ingridient, err := ingridientOld.
		Update().
		SetName(item.Name).
		SetAlcohol(item.Alcohol).
		SetIsDry(item.IsDry).
		SetMeasurmentUnit(ingridient.MeasurmentUnit(item.MeasurmentUnit)).
		Save(c)

	if err != nil {
		ResponseJSON(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	ResponseJSON(c, http.StatusAccepted, resources.IngridientUpdated, ingridient)
}

// @BasePath /api/v1

// AddIngridient godoc
// @Summary            Deletes ingridient
// @Description        Delete ingridient from database by id
// @Tags               Inventory manipulation
// @Produce            application/json
// @Param              itemId query models.ItemModelDelete true "item id"
// @Success            202
// @Router             /ingridient [delete]
func (controller *ItemController) DeleteIngridient(c *gin.Context) {
	jsonData, err := c.GetRawData()
	if err != nil {
		ResponseJSON(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	var itemIdObject models.ItemModelDelete
	json.Unmarshal(jsonData, &itemIdObject)

	ingridientOld, err := controller.Client.Ingridient.
		Query().
		Where(ingridient.ID(itemIdObject.Id)).
		Only(c)

	if err != nil {
		if ingridientOld == nil {
			ResponseJSON(c, http.StatusNotFound, err.Error(), nil)
			return
		}
		ResponseJSON(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	err = controller.Client.Ingridient.
		DeleteOneID(itemIdObject.Id).
		Exec(c)

	if err != nil {
		ResponseJSON(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	ResponseJSON(c, http.StatusAccepted, resources.IngridientDeleted, nil)
}
