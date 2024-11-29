package models

type Book struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ReleaseYear int16  `json:"release_year"`
	Price       int    `json:"price"`
	TotalPage   int16  `json:"total_page"`
	Thickness   string `json:"thickness"`
	CategoryID  int    `json:"category_id"`
}
