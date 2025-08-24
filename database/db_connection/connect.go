package dbconnection

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)


var DB *sql.DB

func getenv(k, d string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return d
}

func OpenDB() (*sql.DB, error) {
	var err error
	if url := os.Getenv("DATABASE_URL"); url != "" {
		DB, err = sql.Open("postgres", url)
		if err != nil {
			return nil, err
		}
		if err := DB.Ping(); err != nil {
			return nil, err
		}
		return DB, nil
	}

	host := getenv("DB_HOST", "localhost")
	port := getenv("DB_PORT", "5432")
	user := getenv("DB_USER", "postgres")
	pass := os.Getenv("DB_PASSWORD")
	name := getenv("DB_NAME", "go_lib")
	ssl := getenv("DB_SSLMODE", "disable") // local "disable"

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		user, pass, host, port, name, ssl)

	DB, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	if err := DB.Ping(); err != nil {
		return nil, err
	}
	return DB, nil
}
