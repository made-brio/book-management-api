package controllers

import (
	"book-management-api/internal/models"
	"book-management-api/internal/service"
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	CategoryService *service.CategoryService
}

func NewCategoryController(service *service.CategoryService) *CategoryController {
	return &CategoryController{CategoryService: service}
}

func (cc *CategoryController) CreateCategory(c *gin.Context) {
	var category models.CreateCategoryRequest

	err := c.BindJSON(&category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = cc.CategoryService.CreateCategory(category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Category created successfully"})
}

func (cc *CategoryController) GetAllCategories(c *gin.Context) {
	var (
		result gin.H
	)

	categories, err := cc.CategoryService.GetAllCategories()

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

func (cc *CategoryController) GetCategoryByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	category, err := cc.CategoryService.GetCategoryByID(id)
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
	id, _ := strconv.Atoi(c.Param("id"))
	err := cc.CategoryService.DeleteCategory(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category has been deleted"})
}
