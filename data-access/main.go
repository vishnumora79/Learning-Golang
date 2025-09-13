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

type Album struct {
	ID int64
	Title string
	Artist string 
	Price float32 
}

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


	albums, err := albumsByArtist("John Coltrane")
	if err != nil {
    	log.Fatal(err)
	}
	fmt.Printf("Albums found: %v\n", albums)
}


func albumsByArtist(name string) ([]Album, error){
	var albums []Album 

	rows, err := db.Query("SELECT * FROM album WHERE artist = $1", name)

	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}

	defer rows.Close()

	for rows.Next() {
		var alb Album 
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}
		albums = append(albums, alb)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	return albums, nil
}