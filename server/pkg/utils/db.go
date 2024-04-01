package utils

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectToDatabase() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Build the connection string
	connectionString := os.Getenv("POSTGRES_URL")

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	DB = db
}
