package main

import (
	"book-management-api/controllers"
	"book-management-api/database"
	"book-management-api/middleware"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	DB  *sql.DB
	err error
)

func main() {
	err = godotenv.Load("config/.env")

	if err != nil {
		fmt.Println("Warning: .env file not found, using default environment variables.")
	}

	psqlInfo := fmt.Sprintf(`host=%s port=%s user=%s password=%s dbname=%s sslmode=disable`,
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	DB, err = sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	defer DB.Close()
	err = DB.Ping()
	if err != nil {
		panic(err)
	}

	database.DBMigrate(DB)

	// Inisialisasi controller
	bookController := controllers.NewBookController(DB)
	categoryController := controllers.NewCategoryController(DB)
	authController := controllers.NewAuthController(DB)

	router := gin.Default()
	// Public routes
	router.POST("/api/users/login", authController.Login)
	router.POST("/api/users/register", authController.Register)

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
	router.Run(":8085")
}
