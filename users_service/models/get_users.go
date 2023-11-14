package models

type GetUsersRequest struct {
}

type GetUsersResponse struct {
	Usernames []string `json:"users"`
}
