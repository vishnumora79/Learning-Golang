package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

var db *sql.DB

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// capture connection properties
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable", 
		os.Getenv("DBUSER"), 
		os.Getenv("DBPASS"), 
		os.Getenv("DBHOST"),
		os.Getenv("DBPORT"),
		os.Getenv("DBNAME"),
	)

	db, err = sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal("Error connecting to the database: ", pingErr)
	}

	fmt.Println("Connected to PostgreSQL!")
}

