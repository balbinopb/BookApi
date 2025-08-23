package main

import (
	dbconnection "book-api/database/db_connection"
	"book-api/database/migration"
	"book-api/initializers"
	"log"
	"net/http"

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
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
