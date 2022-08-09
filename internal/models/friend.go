package models

type Friend struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Sex       int    `json:"sex"`

	PhotoURL string `json:"photo_url"`
}
