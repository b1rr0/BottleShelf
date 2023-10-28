package jwtproc

import (
	"errors"
	"fmt"
	"inventoryService/models"
	"inventoryService/parsistence"
	"inventoryService/parsistence/entities"
	"time"
)

type JwtHandler struct {
	persist parsistence.JwtPersistence
}

func NewJwtHandler(persist parsistence.JwtPersistence) *JwtHandler {
	handler := new(JwtHandler)
	handler.persist = persist
	return handler
}
func (h *JwtHandler) DecodeToken(token string) (UserClaims, error) {
	isValid, err := IsUserTokenValid(token)
	if err != nil {
		return UserClaims{}, err
	}
	if isValid {
		user, err := DecoderTokenToUserClaims(token)
		if !user.isRefresh {
			fmt.Print("------")
			return user, err
		}
	}
	return UserClaims{}, errors.New("time limited")
}

func (h *JwtHandler) RefreshPair(tokenRefresh string) (models.TokenPairs, error) {
	jwt := entities.JwtEntity{Jwt: tokenRefresh}
	if !h.persist.CheckAndRemoveJwt(jwt) {
		return models.TokenPairs{}, errors.New("No Jwt in DB")
	}
	userClaims, err := DecoderTokenToUserClaims(tokenRefresh)
	if err != nil && !userClaims.isRefresh {
		return models.TokenPairs{}, err
	}
	return h.generatePair(userClaims.User)
}

func (h *JwtHandler) CreteNewPair(user models.User) (models.TokenPairs, error) {
	return h.generatePair(user)
}
func (h *JwtHandler) generatePair(user models.User) (models.TokenPairs, error) {
	token, err := BuildUserToken(user, 0*time.Hour, false)
	if err != nil {
		return models.TokenPairs{}, err
	}
	refreshToken, err := BuildUserToken(user, time.Hour*24*30, true)
	if err != nil {
		return models.TokenPairs{}, err
	}
	h.persist.SaveJwt(entities.JwtEntity{Jwt: refreshToken})
	return models.TokenPairs{token, refreshToken}, nil
}
