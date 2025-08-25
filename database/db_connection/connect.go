package dbconnection

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func OpenDB() (*sql.DB, error) {
	url := os.Getenv("DATABASE_URL")
	if url == "" {
		log.Fatal("DATABASE_URL is not set! Please add it in Railway environment variables.")
	}

	var err error
	DB, err = sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	if err := DB.Ping(); err != nil {
		return nil, err
	}

	log.Println("Database connected successfully")
	return DB, nil
}
