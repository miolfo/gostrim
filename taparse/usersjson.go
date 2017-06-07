package taparse

import "time"

type usersJSON struct {
	Total int `json:"_total"`
	Users []struct {
		DisplayName string    `json:"display_name"`
		ID          string    `json:"_id"`
		Name        string    `json:"name"`
		Type        string    `json:"type"`
		Bio         string    `json:"bio"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
		Logo        string    `json:"logo"`
	} `json:"users"`
}
