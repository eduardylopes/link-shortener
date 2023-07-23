package entity

import "time"

type Link struct {
	ID        string    `json:"id"`
	URL       string    `json:"url"`
	Code      string    `json:"code"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
