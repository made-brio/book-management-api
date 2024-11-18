package config

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
	dbURL := os.Getenv("DB_URL") // Railway uses DB_URL as default
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatalf("Database is not reachable: %v", err)
	}
	log.Println("Database connection established.")
	return db
}
