package app

import "time"

type BookModel struct {
	Name        string `db:"name"`
	Description string `db:"description"`
	Quantity    int    `db:"quantity"`
	Author      string `db:"author"`
	CreatedBy   string `db:"created_by"`
}

type ViewBookModel struct {
	ID          int        `db:"id"`
	Name        string     `db:"name"`
	Description *string    `db:"description"`
	Quantity    int        `db:"quantity"`
	Author      string     `db:"author"`
	CreatedAt   time.Time  `db:"created_at"`
	CreatedBy   string     `db:"created_by"`
	UpdateAt    *time.Time `db:"updated_at"`
	UpdateBy    *string    `db:"updated_by"`
}

type TotalRowModel struct {
	Total int `db:"total"`
}

type ViewBookModels []ViewBookModel

type UpdateBookModel struct {
	Name        string `db:"name"`
	Description string `db:"description"`
	Quantity    int    `db:"quantity"`
	Author      string `db:"author"`
	UpdateBy    string `db:"updated_by"`
	Id          int    `db:"id"`
}
