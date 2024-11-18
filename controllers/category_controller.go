package controllers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	DB *sql.DB
}

func NewCategoryController(db *sql.DB) *CategoryController {
	return &CategoryController{DB: db}
}

func (cc *CategoryController) GetAllCategories(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Get all categories"})
}

func (cc *CategoryController) CreateCategory(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"message": "Category created"})
}

func (cc *CategoryController) GetCategoryByID(c *gin.Context) {
	id := c.Param("id")
	var category struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		CreatedAt string `json:"created_at"`
		CreatedBy string `json:"created_by"`
	}

	query := `SELECT id, name, created_at, created_by FROM categories WHERE id = $1`
	err := cc.DB.QueryRow(query, id).Scan(&category.ID, &category.Name, &category.CreatedAt, &category.CreatedBy)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch category details"})
		return
	}

	c.JSON(http.StatusOK, category)
}

func (cc *CategoryController) DeleteCategory(c *gin.Context) {
	id := c.Param("id")

	query := `DELETE FROM categories WHERE id = $1`
	result, err := cc.DB.Exec(query, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete category"})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
}

func (cc *CategoryController) GetBooksByCategory(c *gin.Context) {
	id := c.Param("id")
	rows, err := cc.DB.Query(`
		SELECT id, title, description, image_url, release_year, price, total_page, thickness 
		FROM books 
		WHERE category_id = $1`, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch books by category"})
		return
	}
	defer rows.Close()

	var books []struct {
		ID          int    `json:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
		ImageURL    string `json:"image_url"`
		ReleaseYear int    `json:"release_year"`
		Price       int    `json:"price"`
		TotalPage   int    `json:"total_page"`
		Thickness   string `json:"thickness"`
	}

	for rows.Next() {
		var book struct {
			ID          int    `json:"id"`
			Title       string `json:"title"`
			Description string `json:"description"`
			ImageURL    string `json:"image_url"`
			ReleaseYear int    `json:"release_year"`
			Price       int    `json:"price"`
			TotalPage   int    `json:"total_page"`
			Thickness   string `json:"thickness"`
		}
		if err := rows.Scan(
			&book.ID, &book.Title, &book.Description, &book.ImageURL,
			&book.ReleaseYear, &book.Price, &book.TotalPage, &book.Thickness,
		); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read book data"})
			return
		}
		books = append(books, book)
	}

	c.JSON(http.StatusOK, gin.H{"books": books})
}
