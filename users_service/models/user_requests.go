package models

type CreateUser struct {
	Name     string
	Password string
}

type CheckUser struct {
	Name     string
	Password string
}
