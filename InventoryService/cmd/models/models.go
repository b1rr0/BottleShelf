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
	Name    string  `json:"name"`
	Alcohol float64 `json:"alcohol"`
	IsDry   bool    `json:"isDry"`
}

type ItemModelDelete struct {
	Id uuid.UUID `json:"id"`
}
