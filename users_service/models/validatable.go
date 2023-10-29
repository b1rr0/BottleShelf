package models

type Validatable interface {
	Validate() bool
}
