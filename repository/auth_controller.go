package repository

import (
	"book-management-api/models"
	"database/sql"
	"errors"
)

// GetPasswordByUsername retrieves the hashed password for a given username.
func GetPasswordByUsername(db *sql.DB, username string) (string, error) {
	var password string
	err := db.QueryRow("SELECT password FROM users WHERE username = $1", username).Scan(&password)
	if err == sql.ErrNoRows {
		return "", errors.New("user not found")
	}
	if err != nil {
		return "", err
	}

	return password, nil
}

// CheckUserExists checks if a username already exists in the database.
func CheckUserExists(db *sql.DB, username string) (bool, error) {
	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE username = $1)", username).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

// CreateUser adds a new user to the database.
func CreateUser(db *sql.DB, user models.User) error {
	sqlQuery := `INSERT INTO users (username, password, created_at, created_by)
	VALUES ($1, $2, $3, $4)`
	_, err := db.Exec(sqlQuery, user.Username, user.Password, user.CreatedAt, user.CreatedBy)
	return err
}
