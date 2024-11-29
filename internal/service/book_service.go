package service

import (
	"book-management-api/internal/models"
	"book-management-api/internal/repository"
	"database/sql"
)

type BookService struct {
	DB *sql.DB
}

func NewBookService(db *sql.DB) *BookService {
	return &BookService{DB: db}
}

func (bs *BookService) CreateBook(book models.CreateBookRequest) error {
	return repository.CreateBook(bs.DB, book)
}

func (bs *BookService) GetAllBooks() ([]models.Book, error) {
	return repository.GetAllBooks(bs.DB)
}

func (bs *BookService) GetBookByID(id int) (models.Book, error) {
	return repository.GetBookByID(bs.DB, id)
}

func (bs *BookService) GetBooksByCategory(id int) ([]models.Book, error) {
	return repository.GetBooksByCategory(bs.DB, id)
}

func (bs *BookService) DeleteBook(id int) error {
	return repository.DeleteBook(bs.DB, id)
}
