package main

import (
	"book-api/config/initializers"
	dbconnection "book-api/database/db_connection"
	"book-api/database/migration"
	"book-api/router"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
}

func main() {

	db, err := dbconnection.OpenDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	migration.DBMigrate(db)

	r := gin.Default()

	router.Routes(r)


	port := os.Getenv("PGPORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("\nServer running di: http://localhost:%s ...\n", port)
	r.Run(":" + port)
}
