package controllers

import (
	"book-management-api/utils"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	DB *sql.DB
}

func NewAuthController(db *sql.DB) *AuthController {
	return &AuthController{DB: db}
}

func (ac *AuthController) Login(c *gin.Context) {
	var user struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// Bind JSON input to struct
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Check user in database
	var storedPassword string
	err := ac.DB.QueryRow("SELECT password FROM users WHERE username = $1", user.Username).Scan(&storedPassword)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Validate password
	if !utils.CheckPasswordHash(user.Password, storedPassword) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate JWT token
	token, err := generateJWT(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Return the JWT token
	c.JSON(http.StatusOK, gin.H{"token": token})
}

// Function to generate JWT token
func generateJWT(username string) (string, error) {
	// Define JWT secret key (this should be kept secure, not hard-coded in production)
	secretKey := []byte("your_secret_key_here")

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Set expiration time (24 hours)
	})

	// Generate signed token
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", fmt.Errorf("failed to sign the token: %w", err)
	}

	return signedToken, nil
}
