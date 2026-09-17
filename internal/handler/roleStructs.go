package handler

type getRoleResponse struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	Permissions []string `json:"permissions"`
}
