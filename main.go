package main

import (
	"book-api/config/initializers"
	dbconnection "book-api/database/db_connection"
	"book-api/database/migration"
	"book-api/router"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	initializers.LoadEnv()

	db, err := dbconnection.OpenDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	migration.DBMigrate(db)

	r := gin.Default()

	router.Routes(r)


	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
