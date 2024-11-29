package repository

import (
	"book-management-api/internal/models"
	"database/sql"
	"errors"
	"log"
)

// GetPasswordByUsername retrieves the hashed password for a given username.
func GetPasswordByUsername(db *sql.DB, username string) (string, error) {
	var password string
	err := db.QueryRow("SELECT password FROM user_account WHERE username = $1", username).Scan(&password)
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
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM user_account WHERE username = $1)", username).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

// CreateUser inserts a new user into the database
func CreateUser(db *sql.DB, user models.UserAccount) (err error) {
	sqlQuery := `INSERT INTO user_account (username, password) VALUES ($1, $2)`
	_, err = db.Exec(sqlQuery, user.Username, user.Password)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		return err
	}

	return err
}
