package controllers

import (
	"book-management-api/models"
	"book-management-api/repository"
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookController struct {
	DB *sql.DB
}

func NewBookController(db *sql.DB) *BookController {
	return &BookController{DB: db}
}

func (bc *BookController) GetAllBooks(c *gin.Context) {
	var (
		result gin.H
	)

	book, err := repository.GetAllBooks(bc.DB)

	if err != nil {
		result = gin.H{
			"result": err.Error(),
		}
	} else {
		result = gin.H{
			"result": book,
		}
	}

	c.JSON(http.StatusOK, result)
}

func (bc *BookController) CreateBook(c *gin.Context) {
	var book models.Book

	err := c.BindJSON(&book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = repository.CreateBook(bc.DB, book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, book)
}

func (bc *BookController) GetBookByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	book, err := repository.GetBookByID(bc.DB, id)
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
	var book models.Book
	id, _ := strconv.Atoi(c.Param("id"))
	book.ID = id
	err := repository.DeleteBook(bc.DB, book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, book)
}
