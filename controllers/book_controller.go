package controllers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookController struct {
	DB *sql.DB
}

func NewBookController(db *sql.DB) *BookController {
	return &BookController{DB: db}
}

func (bc *BookController) GetAllBooks(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Get all books"})
}

func (bc *BookController) CreateBook(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"message": "Book created"})
}

func (bc *BookController) GetBookByID(c *gin.Context) {
	id := c.Param("id")
	var book struct {
		ID          int    `json:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
		ImageURL    string `json:"image_url"`
		ReleaseYear int    `json:"release_year"`
		Price       int    `json:"price"`
		TotalPage   int    `json:"total_page"`
		Thickness   string `json:"thickness"`
		CategoryID  int    `json:"category_id"`
	}

	query := `SELECT id, title, description, image_url, release_year, price, total_page, thickness, category_id FROM books WHERE id = $1`
	err := bc.DB.QueryRow(query, id).Scan(
		&book.ID,
		&book.Title,
		&book.Description,
		&book.ImageURL,
		&book.ReleaseYear,
		&book.Price,
		&book.TotalPage,
		&book.Thickness,
		&book.CategoryID,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch book details"})
		return
	}

	c.JSON(http.StatusOK, book)
}

func (bc *BookController) DeleteBook(c *gin.Context) {
	id := c.Param("id")

	query := `DELETE FROM books WHERE id = $1`
	result, err := bc.DB.Exec(query, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete book"})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}
