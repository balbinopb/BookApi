package dbconnection

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func OpenDB() (*sql.DB, error) {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("PGHOST"),
		os.Getenv("PGPORT"),
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
		os.Getenv("PGDATABASE"),
	)

	DB, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("FAILED TO OPEN DATABASE:", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatal("FAILED TO PING DATABASE:", err)
	}

	log.Println("DATABASE IS CONNECTED SUCCESFULLY!!!")
	return DB, nil
}
