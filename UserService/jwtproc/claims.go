package jwtproc

import (
	"github.com/golang-jwt/jwt/v5"
	"inventoryService/models"
)

type UserClaims struct {
	jwt.RegisteredClaims
	models.User
	isRefresh bool
}
