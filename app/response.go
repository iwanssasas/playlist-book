package app

type BookResponse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Quantity    int    `josn:"quantity"`
	Author      string `json:"author"`
	CreatedBy   string `json:"created_by"`
}
