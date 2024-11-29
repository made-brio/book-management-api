package models

//general

//user
type UserLoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

//book
type CreateBookRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	ReleaseYear int16  `json:"release_year" binding:"required"`
	Price       int    `json:"price" binding:"required"`
	TotalPage   int16  `json:"total_page" binding:"required"`
	CategoryID  int    `json:"category_id" binding:"required"`
}

//category
type CreateCategoryRequest struct {
	Name string `json:"name" binding:"required"`
}
