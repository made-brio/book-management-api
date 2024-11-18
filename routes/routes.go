package routes

import (
	"book-management-api/controllers"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, db *sql.DB) {
	// Categories
	categoryController := controllers.NewCategoryController(db)
	router.GET("/api/categories", categoryController.GetAllCategories)
	router.POST("/api/categories", categoryController.CreateCategory)
	router.GET("/api/categories/:id", categoryController.GetCategoryByID)
	router.DELETE("/api/categories/:id", categoryController.DeleteCategory)
	router.GET("/api/categories/:id/books", categoryController.GetBooksByCategory)

	// Books
	bookController := controllers.NewBookController(db)
	router.GET("/api/books", bookController.GetAllBooks)
	router.POST("/api/books", bookController.CreateBook)
	router.GET("/api/books/:id", bookController.GetBookByID)
	router.DELETE("/api/books/:id", bookController.DeleteBook)
}
