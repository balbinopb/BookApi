package dbconnection

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func getenv(k, d string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return d
}

func OpenDB() (*sql.DB, error) {

	if url := os.Getenv("DATABASE_URL"); url != "" {
		db, err := sql.Open("postgres", url)
		if err != nil {
			return nil, err
		}
		if err := db.Ping(); err != nil {
			return nil, err
		}
		return db, nil
	}

	host := getenv("DB_HOST", "localhost")
	port := getenv("DB_PORT", "5432")
	user := getenv("DB_USER", "postgres")
	pass := os.Getenv("DB_PASSWORD")
	name := getenv("DB_NAME", "go_lib")
	ssl := getenv("DB_SSLMODE", "disable") // local "disable"

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		user, pass, host, port, name, ssl)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
