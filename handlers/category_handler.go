package handlers

import (
	dbconnection "book-api/database/db_connection"
	"book-api/models"
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get all
// GET /api/categories
func GetCategories(c *gin.Context) {
	rows, err := dbconnection.DB.Query("SELECT id, name, created_at, created_by, modified_at, modified_by FROM categories")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB error"})
		return
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var cat models.Category
		if err := rows.Scan(&cat.ID, &cat.Name, &cat.CreatedAt, &cat.CreatedBy, &cat.ModifiedAt, &cat.ModifiedBy); err == nil {
			categories = append(categories, cat)
		}
	}
	c.JSON(http.StatusOK, categories)
}

// Create new categorie
// POST /api/categories
func CreateCategory(c *gin.Context) {
	var cat models.Category
	if err := c.ShouldBindJSON(&cat); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := dbconnection.DB.QueryRow(
		"INSERT INTO categories (name, created_at, created_by) VALUES ($1, NOW(), 'admin') RETURNING id, name, created_at, created_by",
		cat.Name,
	).Scan(&cat.ID, &cat.Name, &cat.CreatedAt, &cat.CreatedBy)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Insert failed"})
		return
	}

	c.JSON(http.StatusCreated, cat)
}

// Get by id
// GET /api/categories/:id
func GetCategoryByID(c *gin.Context) {
	id := c.Param("id")
	var cat models.Category
	err := dbconnection.DB.QueryRow("SELECT id, name, created_at, created_by, modified_at, modified_by FROM categories WHERE id=$1", id).
		Scan(&cat.ID, &cat.Name, &cat.CreatedAt, &cat.CreatedBy, &cat.ModifiedAt, &cat.ModifiedBy)

	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}
	if err != nil {
		log.Printf("============Get by id error: %v=============", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB error"})
		return
	}

	c.JSON(http.StatusOK, cat)
}

// DELETE by id
// DELETE /api/categories/:id
func DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	res, err := dbconnection.DB.Exec("DELETE FROM categories WHERE id=$1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Delete failed"})
		return
	}
	count, _ := res.RowsAffected()
	if count == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Deleted Success"})
}


// GET /api/categories/:id/books
func GetBooksByCategory(c *gin.Context) {
	categoryID := c.Param("id")

	rows, err := dbconnection.DB.Query(`
		SELECT id, title, description, image_url, release_year, price, total_page, thickness, category_id, created_at, created_by, modified_at, modified_by 
		FROM books WHERE category_id = $1`, categoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB query error"})
		return
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var b models.Book
		if scanErr := rows.Scan(
			&b.ID, &b.Title, &b.Description, &b.ImageURL,
			&b.ReleaseYear, &b.Price, &b.TotalPage, &b.Thickness,
			&b.CategoryID, &b.CreatedAt, &b.CreatedBy,
			&b.ModifiedAt, &b.ModifiedBy,
		); scanErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Row scan error"})
			return
		}
		books = append(books, b)
	}

	// Important: check rows.Err() after iteration
	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Iteration error"})
		return
	}

	if len(books) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No books found for this category"})
		return
	}

	c.JSON(http.StatusOK, books)
}

