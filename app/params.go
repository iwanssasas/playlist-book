package app

type BookParams struct {
	Name        string `form:"name"`
	Description string `form:"description"`
	Quantity    int    `form:"quantity"`
	Author      string `form:"author"`
}

type QueryParams struct {
	Page string `form:"page" binding:require`
}
