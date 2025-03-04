package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUser := os.Getenv("POSTGRES_USER") // Default: "postgres"
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbHost := "localhost"
	dbPort := "5433"
	dbName := "brightpixel"
	dbSSL := "disable"

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", dbUser, dbPassword, dbHost, dbPort, dbName, dbSSL)

	db, err := sqlx.Connect("postgres", dbURL)
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	fmt.Println("Connected to PostgreSQL!")
	DB = db
}
