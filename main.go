package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	"book-management-api/controllers"
	"book-management-api/middleware"
)

func main() {
	// Koneksi ke database
	db, err := sql.Open("postgres", "your_connection_string_here")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// Inisialisasi controller
	bookController := controllers.NewBookController(db)
	categoryController := controllers.NewCategoryController(db)

	// Inisialisasi router Gin
	router := gin.Default()

	// Public routes
	router.POST("/api/users/login", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Login endpoint"})
	})

	// Protected routes dengan middleware JWTAuth
	protected := router.Group("/api", middleware.JWTAuth())
	{
		// Rute untuk buku
		protected.GET("/books", bookController.GetAllBooks)
		protected.POST("/books", bookController.CreateBook)
		protected.GET("/books/:id", bookController.GetBookByID)
		protected.DELETE("/books/:id", bookController.DeleteBook)

		// Rute untuk kategori
		protected.GET("/categories", categoryController.GetAllCategories)
		protected.POST("/categories", categoryController.CreateCategory)
		protected.GET("/categories/:id", categoryController.GetCategoryByID)
		protected.DELETE("/categories/:id", categoryController.DeleteCategory)
		protected.GET("/categories/:id/books", categoryController.GetBooksByCategory)
	}

	// Menjalankan server
	router.Run(":8080")
}
