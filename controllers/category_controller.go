package controllers

import (
	"book-management-api/models"
	"book-management-api/repository"
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	DB *sql.DB
}

func NewCategoryController(db *sql.DB) *CategoryController {
	return &CategoryController{DB: db}
}

func (cc *CategoryController) GetAllCategories(c *gin.Context) {
	var (
		result gin.H
	)

	categories, err := repository.GetAllCategories(cc.DB)

	if err != nil {
		result = gin.H{
			"result": err.Error(),
		}
	} else {
		result = gin.H{
			"result": categories,
		}
	}

	c.JSON(http.StatusOK, result)
}

func (cc *CategoryController) CreateCategory(c *gin.Context) {
	var category models.Category

	err := c.BindJSON(&category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = repository.CreateCategory(cc.DB, category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, category)
}

func (cc *CategoryController) GetCategoryByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	category, err := repository.GetCategoryByID(cc.DB, id)
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
	var category models.Category
	id, _ := strconv.Atoi(c.Param("id"))
	category.ID = id
	err := repository.DeleteCategory(cc.DB, category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, category)
}

func (cc *CategoryController) GetBooksByCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	books, err := repository.GetBooksByCategory(cc.DB, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"books": books})
}
