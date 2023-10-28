package parsistence

import (
	"inventoryService/parsistence/entities"
)

type LocalPersister struct {
	jwts []entities.JwtEntity
}

func (p *LocalPersister) CheckAndRemoveJwt(jwt entities.JwtEntity) bool {
	for i, j := range p.jwts {
		if j.Jwt == jwt.Jwt {
			p.jwts[i] = p.jwts[len(p.jwts)-1]
			p.jwts[len(p.jwts)-1] = entities.JwtEntity{}
			p.jwts = p.jwts[:len(p.jwts)-1]
			return true
		}
	}
	return false
}

func (p *LocalPersister) SaveJwt(jwt entities.JwtEntity) {
	p.jwts = append(p.jwts, jwt)
}
