package routes

import (
	"book-management-api/internal/controllers"
	"book-management-api/internal/middleware"
	"book-management-api/internal/service"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, db *sql.DB) {

	// Initialize Services
	bookService := service.NewBookService(db)
	categoryService := service.NewCategoryService(db)
	authService := service.NewAuthService(db)

	// Authentication
	authController := controllers.NewAuthController(authService)
	authRoutes := router.Group("/api/users")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.CreateUser)
	}

	//categories
	categoryController := controllers.NewCategoryController(categoryService)
	categoriesRoutes := router.Group("/api/categories", middleware.JWTAuth())
	{
		categoriesRoutes.GET("/", categoryController.GetAllCategories)
		categoriesRoutes.POST("/", categoryController.CreateCategory)
		categoriesRoutes.GET("/:id", categoryController.GetCategoryByID)
		categoriesRoutes.DELETE("/:id", categoryController.DeleteCategory)
	}

	// Books
	bookController := controllers.NewBookController(bookService)
	booksRoutes := router.Group("/api/books", middleware.JWTAuth())
	{
		booksRoutes.GET("/", bookController.GetAllBooks)
		booksRoutes.POST("/", bookController.CreateBook)
		booksRoutes.GET("/categories/:id", bookController.GetBooksByCategory)
		booksRoutes.GET("/:id", bookController.GetBookByID)
		booksRoutes.DELETE("/:id", bookController.DeleteBook)
	}

}
