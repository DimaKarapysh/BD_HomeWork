package db

import (
	"database/sql"
	"fmt"
	"github.com/pressly/goose"
	"log"
	"os"
)

var (
	DB *sql.DB
)

func connectionString() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))
}

// InitDatabase For Vanilla SQL
func InitDatabase() (*sql.DB, error) {
	db, err := sql.Open(os.Getenv("DB_DRIVER"), connectionString())
	if err != nil {
		return nil, err
	}
	DB = db
	log.Println("Connected to database")
	return db, nil
}

func RunMigrations() error {
	err := goose.Up(DB, "./migrations")
	if err != nil {
		return err
	}
	log.Println("Migrations applied successfully")
	return nil
}
