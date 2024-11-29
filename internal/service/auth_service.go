package service

import (
	"book-management-api/internal/models"
	"book-management-api/internal/repository"
	"database/sql"
)

// AuthService defines the service layer for Authorization user
type AuthService struct {
	DB *sql.DB
}

func NewAuthService(db *sql.DB) *AuthService {
	return &AuthService{DB: db}
}

// CreateUser creates a new user
func (as *AuthService) CreateUser(user models.UserAccount) error {
	return repository.CreateUser(as.DB, user)
}

func (as *AuthService) GetPasswordByUsername(user string) (string, error) {
	return repository.GetPasswordByUsername(as.DB, user)
}

func (as *AuthService) CheckUserExists(user string) (bool, error) {
	return repository.CheckUserExists(as.DB, user)
}
