package router

import (
	"book-api/handlers"
	middleware "book-api/middleware"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	router := r.Group("/api", middleware.BasicAuthMiddleware())


	//categories
	router.GET("/categories", handlers.GetCategories)
	router.POST("/categories", handlers.CreateCategory)
	router.GET("/categories/:id", handlers.GetCategoryByID)
	router.DELETE("/categories/:id", handlers.DeleteCategory)
	router.GET("/categories/:id/books", handlers.GetBooksByCategory)

	// books
	router.GET("/books", handlers.GetBooks)
	router.POST("/books", handlers.CreateBook)
	router.GET("/books/:id", handlers.GetBookByID)
	router.DELETE("/books/:id", handlers.DeleteBook)
	router.PUT("/books/:id", handlers.UpdateBook)
}
