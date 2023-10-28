package jwtproc

import (
	"authservice/models"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// ToDo: Extract or generate
const jwtSecret = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJpc3MiOiJ0ZXN0IiwiYXVkIjoic2luZ2xlIn0.QAWg1vGvnqRuCFTMcPkjZljXHh8U3L_qUjszOtQbeaA"

func BuildUserToken(user models.User, duration time.Duration, isRefresh bool) (string, error) {
	claims := &UserClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
		}, user, isRefresh,
	}
	return buildToken(claims)
}

func buildToken(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecret))
}

func DecoderTokenToUserClaims(tokenString string) (UserClaims, error) {
	claims, err := decodeToken(tokenString, &UserClaims{})
	cl := claims.(*UserClaims)
	return *cl, err
}

func IsUserTokenValid(tokenString string) (bool, error) {
	claims, err := decodeToken(tokenString, &UserClaims{})
	if err != nil {
		return false, err
	}
	cl := claims.(*UserClaims)
	return cl.ExpiresAt.Before(time.Now()), nil
}

func decodeToken(tokenString string, cl jwt.Claims) (jwt.Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, cl, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	return token.Claims, err
}
