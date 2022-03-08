package app

type BookModel struct {
	Name        string `db:"name"`
	Description string `db:"description"`
	Quantity    int    `db:"quantity"`
	Author      string `db:"author"`
	CreatedBy   string `db:"created_by"`
}
