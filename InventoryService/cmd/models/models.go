package models

import (
	"github.com/google/uuid"
)

type Unit int

type ItemModel struct {
	Id             uuid.UUID `json:"id"`
	Name           string    `json:"name"`
	Alcohol        float64   `json:"alcohol"`
	IsDry          bool      `json:"isDry"`
	MeasurmentUnit string    `json:"measurmentUnit"`
}

type ItemModelCreate struct {
	Name           string  `json:"name"`
	Alcohol        float64 `json:"alcohol"`
	IsDry          bool    `json:"isDry"`
	MeasurmentUnit string  `json:"measurmentUnit"`
}

type ItemModelFilters struct {
	Name       string  `form:"name"`
	AlcoholMin float64 `form:"alcoholmin"`
	AlcoholMax float64 `form:"alcoholmax"`
	IsDry      bool    `form:"isDry"`
}

type ItemModelDelete struct {
	Id uuid.UUID `json:"id"`
}
