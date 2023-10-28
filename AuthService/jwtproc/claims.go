package jwtproc

import (
	"authservice/models"
	"github.com/golang-jwt/jwt/v5"
)

type UserClaims struct {
	jwt.RegisteredClaims
	models.User
	isRefresh bool
}
