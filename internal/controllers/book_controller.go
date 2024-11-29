package controllers

import (
	"book-management-api/internal/models"
	"book-management-api/internal/service"
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookController struct {
	BookService *service.BookService
}

func NewBookController(service *service.BookService) *BookController {
	return &BookController{BookService: service}
}

func (bc *BookController) CreateBook(c *gin.Context) {
	var book models.CreateBookRequest
	err := c.BindJSON(&book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = bc.BookService.CreateBook(book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Book created successfully"})
}

func (bc *BookController) GetAllBooks(c *gin.Context) {
	var (
		result gin.H
	)

	books, err := bc.BookService.GetAllBooks()

	if err != nil {
		result = gin.H{
			"result": err.Error(),
		}
	} else {
		result = gin.H{
			"result": books,
		}
	}

	c.JSON(http.StatusOK, result)
}

func (bc *BookController) GetBookByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	result, err := bc.BookService.GetBookByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch book details"})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (bc *BookController) GetBooksByCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	result, err := bc.BookService.GetBooksByCategory(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": result})
}

func (bc *BookController) DeleteBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	err := bc.BookService.DeleteBook(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book has been deleted"})
}
