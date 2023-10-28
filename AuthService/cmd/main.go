package main

import (
	"authservice/jwtproc"
	"authservice/models"
	"authservice/parsistence"
)

func main() {
	parsi := parsistence.LocalPersister{}
	handler := jwtproc.NewJwtHandler(&parsi)

	user := models.User{
		UserId:       "aksjdasd",
		Role:         "kajsdas",
		Organisation: "kasjd",
	}
	pair, _ := handler.CreteNewPair(user)
	handler.DecodeToken(pair.Jwt)
	handler.RefreshPair(pair.RefreshJwt)
}
