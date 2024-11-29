package service

import (
	"book-management-api/internal/models"
	"book-management-api/internal/repository"
	"database/sql"
)

type CategoryService struct {
	DB *sql.DB
}

func NewCategoryService(db *sql.DB) *CategoryService {
	return &CategoryService{DB: db}
}

func (cs *CategoryService) CreateCategory(payload models.CreateCategoryRequest) error {
	return repository.CreateCategory(cs.DB, payload)
}

func (cs *CategoryService) GetAllCategories() ([]models.Category, error) {
	return repository.GetAllCategories(cs.DB)
}

func (cs *CategoryService) GetCategoryByID(id int) (models.Category, error) {
	return repository.GetCategoryByID(cs.DB, id)
}

func (cs *CategoryService) DeleteCategory(id int) error {
	return repository.DeleteCategory(cs.DB, id)
}
