package controllers

import (
	"encoding/json"
	"fmt"
	"inventoryService/m/v2/cmd/models"
	"inventoryService/m/v2/cmd/resources"
	"inventoryService/m/v2/ent"
	"inventoryService/m/v2/ent/ingridient"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ItemController struct {
	Client *ent.Client
}

func (controller *ItemController) ValidateNewIngridient(c *gin.Context, item models.ItemModelCreate) (exists bool, valid bool, errorMessage []string) {
	exists, err := controller.Client.Ingridient.
		Query().
		Where(ingridient.Name(item.Name)).
		Exist(c)

	valid = true

	if err != nil {
		resources.ResponseJSON(c, http.StatusServiceUnavailable, err.Error(), nil)
		return
	}

	if item.Alcohol < 0 || item.Alcohol > 1 {
		errorMessage = append(errorMessage, "Alcohol")
		valid = false
		return
	}

	if !(item.MeasurmentUnit == "pcs" || item.MeasurmentUnit == "g" || item.MeasurmentUnit == "ml" || item.MeasurmentUnit == "") {
		errorMessage = append(errorMessage, "MeasurmentUnit")
		valid = false
		return
	}
	return
}

// @BasePath /api/v1

// GetIngridientsList godoc
// @Summary        Gets list of all ingridients
// @Description    Get complete list of all ingridients availible for user
// @Tags           Inventory manipulation
// @Produce        application/json
// @Success        200 {object} models.ItemModel
// @Router         /inventory [get]
func (controller *ItemController) GetIngridientsList(c *gin.Context) {
	ingridients, err := controller.Client.Ingridient.
		Query().
		All(c)

	if err != nil {
		resources.ResponseJSON(c, http.StatusServiceUnavailable, err.Error(), nil)
		return
	}

	resources.ResponseJSON(c, http.StatusOK, "", ingridients)
}

// @BasePath /api/v1

// GetIngridientsList godoc
// @Summary        Gets ingridient
// @Description    Get ingridient with specified id
// @Tags           Inventory manipulation
// @Produce        application/json
// @Param          Id query models.ItemModelId true "Id of item"
// @Success        200 {object} models.ItemModel
// @Failure        400 {string} string
// @Failure        404 {string} string
// @Router         /ingridient [get]
func (controller *ItemController) GetIngridientById(c *gin.Context) {
	idString, exists := c.GetQuery("id")

	if !exists {
		resources.ResponseJSON(c, http.StatusBadRequest, resources.IdNotFound, nil)
		return
	}

	var id, err = uuid.Parse(idString)

	if err != nil {
		resources.ResponseJSON(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	ingridient, err := controller.Client.Ingridient.
		Query().
		Where(ingridient.ID(id)).
		Only(c)

	if err != nil {
		if ent.IsNotFound(err) {
			resources.ResponseJSON(c, http.StatusNotFound, resources.IngridientNotFound, nil)
			return
		}
		resources.ResponseJSON(c, http.StatusServiceUnavailable, err.Error(), nil)
		return
	}

	resources.ResponseJSON(c, http.StatusOK, "", ingridient)
}

// @BasePath /api/v1

// GetIngridientsList godoc
// @Summary         Gets list ingridients by filter
// @Description     Get list of ingridients filtering by it's name and/or parameters
// @Tags            Inventory manipulation
// @Produce         application/json
// @Param           item query models.ItemModelFilters true "Item"
// @Success         200 {object} models.ItemModel
// @Failure         400 {string} string
// @Router          /inventory/search [get]
func (controller *ItemController) GetIngridientsByFilter(c *gin.Context) {
	var filters models.ItemModelFilters
	filters.AlcoholMax = 1

	err := c.BindQuery(&filters)

	if err != nil {
		resources.ResponseJSON(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if filters.AlcoholMin < 0 || filters.AlcoholMax > 1 || filters.AlcoholMax < 0 || filters.AlcoholMin > 1 {
		resources.ResponseJSON(c, http.StatusBadRequest, resources.InadequateAlcohol, nil)
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
		resources.ResponseJSON(c, http.StatusServiceUnavailable, err.Error(), nil)
		return
	}

	resources.ResponseJSON(c, http.StatusOK, "", ingridients)
}

// @BasePath /api/v1

// AddIngridient godoc
// @Summary           Adds new ingridient
// @Description       Add new ingridient to database
// @Tags              Inventory manipulation
// @Accept            json
// @Produce           application/json
// @Param             item body models.ItemModelCreate true "Item data"
// @Success           201 {object} models.ItemModel
// @Failure           400 {string} string
// @Router            /ingridient [post]
func (controller *ItemController) AddIngridient(c *gin.Context) {
	jsonData, err := c.GetRawData()
	if err != nil {
		resources.ResponseJSON(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	var item models.ItemModelCreate
	err = json.Unmarshal(jsonData, &item)

	if err != nil {
		resources.ResponseJSON(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	exists, valid, errorMessage := controller.ValidateNewIngridient(c, item)

	if exists {
		resources.ResponseJSON(c, http.StatusBadRequest, resources.AlreadyExists, nil)
		return
	}

	if !valid {
		resources.ResponseJSON(c, http.StatusBadRequest, fmt.Sprintf(resources.WrongParameter, errorMessage[0]), nil)
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
		resources.ResponseJSON(c, http.StatusServiceUnavailable, err.Error(), nil)
		return
	}

	resources.ResponseJSON(c, http.StatusCreated, resources.IngridientAdded, ingridient)
}

// @BasePath /api/v1

// AddIngridient godoc
// @Summary           Changes ingridient information
// @Description       Change ingridient in the database by id. All fields are required, otherwise default value will be used.
// @Tags              Inventory manipulation
// @Accept            json
// @Produce           application/json
// @Param             item body models.ItemModel true "Item id and it's new data"
// @Success           202 {object} models.ItemModel
// @Failure           400 {string} string
// @Failure           404 {string} string
// @Router            /ingridient [put]
func (controller *ItemController) ChangeIngridient(c *gin.Context) {
	jsonData, err := c.GetRawData()
	if err != nil {
		resources.ResponseJSON(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	var item models.ItemModel
	err = json.Unmarshal(jsonData, &item)

	if err != nil {
		resources.ResponseJSON(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	ingridientOld, err := controller.Client.Ingridient.
		Query().
		Where(ingridient.ID(item.Id)).
		Only(c)

	if err != nil {
		if ent.IsNotFound(err) {
			resources.ResponseJSON(c, http.StatusNotFound, resources.IngridientNotFound, nil)
			return
		}
		resources.ResponseJSON(c, http.StatusServiceUnavailable, err.Error(), nil)
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
		resources.ResponseJSON(c, http.StatusServiceUnavailable, err.Error(), nil)
		return
	}

	resources.ResponseJSON(c, http.StatusAccepted, resources.IngridientUpdated, ingridient)
}

// @BasePath /api/v1

// AddIngridient godoc
// @Summary            Deletes ingridient
// @Description        Delete ingridient from database by id
// @Tags               Inventory manipulation
// @Produce            application/json
// @Param              itemId query models.ItemModelId true "item id"
// @Success            202 {string} string
// @Failure            400 {string} string
// @Failure            404 {string} string
// @Router             /ingridient [delete]
func (controller *ItemController) DeleteIngridient(c *gin.Context) {
	idString, exists := c.GetQuery("id")

	if !exists {
		resources.ResponseJSON(c, http.StatusBadRequest, resources.IdNotFound, nil)
		return
	}

	var id, err = uuid.Parse(idString)

	if err != nil {
		resources.ResponseJSON(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	_, err = controller.Client.Ingridient.
		Query().
		Where(ingridient.ID(id)).
		Only(c)

	if err != nil {
		if ent.IsNotFound(err) {
			resources.ResponseJSON(c, http.StatusNotFound, resources.IngridientNotFound, nil)
			return
		}
		resources.ResponseJSON(c, http.StatusServiceUnavailable, err.Error(), nil)
		return
	}

	err = controller.Client.Ingridient.
		DeleteOneID(id).
		Exec(c)

	if err != nil {
		resources.ResponseJSON(c, http.StatusServiceUnavailable, err.Error(), nil)
		return
	}

	resources.ResponseJSON(c, http.StatusAccepted, resources.IngridientDeleted, nil)
}
