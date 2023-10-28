package parsistence

import "inventoryService/parsistence/entities"

type JwtPersistence interface {
	SaveJwt(jwt entities.JwtEntity)
	CheckAndRemoveJwt(jwt entities.JwtEntity) bool
}
