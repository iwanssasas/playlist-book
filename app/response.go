package app

import "time"

type BookResponse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Quantity    int    `josn:"quantity"`
	Author      string `json:"author"`
	CreatedBy   string `json:"created_by"`
}

type SelectBookResponse struct {
	ID          int        `json:"id"`
	Name        string     `json:"name"`
	Description *string    `json:"description"`
	Quantity    int        `json:"quantity"`
	Author      string     `json:"author"`
	CreatedAt   time.Time  `json:"created_at"`
	CreatedBy   string     `json:"created_by"`
	UpdateAt    *time.Time `json:"updated_at"`
	UpdateBy    *string    `json:"updated_by"`
}

type SelectBookResponses []SelectBookResponse
