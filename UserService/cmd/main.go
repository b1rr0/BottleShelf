package main

import (
	"fmt"
	"inventoryService/jwtproc"
	"inventoryService/models"
	"inventoryService/parsistence"
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
	fmt.Println(handler.RefreshPair(pair.RefreshJwt))

}
