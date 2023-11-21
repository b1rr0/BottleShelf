package models

type GetOrgsRequest struct {
}

type GetOrgsResponse struct {
	Orgnames []string `json:"orgs"`
}
