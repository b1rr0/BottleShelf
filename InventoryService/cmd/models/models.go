package models

import (
	"github.com/google/uuid"
)

const (
	ml  Unit = 0
	g   Unit = 1
	pcs Unit = 2
)

type Unit int

type ItemModel struct {
	Id             uuid.UUID
	Name           string
	Alcohol        float64
	IsDry          bool
	MeasurmentUnit Unit
}
