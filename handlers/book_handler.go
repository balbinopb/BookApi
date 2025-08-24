package handlers

import (
	dbconnection "book-api/database/db_connection"
	"book-api/models"
	"book-api/utils"
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Get all books
// GET /api/books
func GetBooks(c *gin.Context) {
	rows, err := dbconnection.DB.Query(`SELECT id, title, description, image_url, release_year, price, total_page, thickness, category_id, created_at, created_by, modified_at, modified_by FROM books`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB error"})
		return
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var b models.Book
		if err := rows.Scan(&b.ID, &b.Title, &b.Description, &b.ImageURL, &b.ReleaseYear, &b.Price, &b.TotalPage, &b.Thickness, &b.CategoryID, &b.CreatedAt, &b.CreatedBy, &b.ModifiedAt, &b.ModifiedBy); err == nil {
			books = append(books, b)
		}
	}
	c.JSON(http.StatusOK, books)
}


// Post book
// POST /api/books
func CreateBook(c *gin.Context) {
	var b models.Book
	if err := c.ShouldBindJSON(&b); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Validation release_year
	if b.ReleaseYear < 1980 || b.ReleaseYear > 2024 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "release_year must be between 1980 and 2024"})
		return
	}

	// Set thickness
	b.Thickness = utils.GetThickness(b.TotalPage)

	// Insert to DB
	err := dbconnection.DB.QueryRow(
		`INSERT INTO books (title, description, image_url, release_year, price, total_page, thickness, category_id, created_at, created_by) 
		 VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10) RETURNING id`,
		b.Title, b.Description, b.ImageURL, b.ReleaseYear, b.Price, b.TotalPage, b.Thickness, b.CategoryID, time.Now(), "admin",
	).Scan(&b.ID)

	if err != nil {
		log.Printf("============Post book error: %v=============", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Insert failed"})
		return
	}

	c.JSON(http.StatusCreated, b)
}

// GET /api/books/:id
func GetBookByID(c *gin.Context) {
	id := c.Param("id")
	var b models.Book
	err := dbconnection.DB.QueryRow(`SELECT id, title, description, image_url, release_year, price, total_page, thickness, category_id, created_at, created_by, modified_at, modified_by 
		FROM books WHERE id=$1`, id).
		Scan(&b.ID, &b.Title, &b.Description, &b.ImageURL, &b.ReleaseYear, &b.Price, &b.TotalPage, &b.Thickness, &b.CategoryID, &b.CreatedAt, &b.CreatedBy, &b.ModifiedAt, &b.ModifiedBy)

	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB error"})
		return
	}

	c.JSON(http.StatusOK, b)
}

// DELETE /api/books/:id
func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	res, err := dbconnection.DB.Exec("DELETE FROM books WHERE id=$1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Delete failed"})
		return
	}
	count, _ := res.RowsAffected()
	if count == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Deleted successfull"})
}

// PUT /api/books/:id  (optional update endpoint)
func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var b models.Book
	if err := c.ShouldBindJSON(&b); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Validation release_year
	if b.ReleaseYear < 1980 || b.ReleaseYear > 2024 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "release_year must be between 1980 and 2024"})
		return
	}

	// Update thickness
	b.Thickness = utils.GetThickness(b.TotalPage)

	res, err := dbconnection.DB.Exec(`UPDATE books SET title=$1, description=$2, image_url=$3, release_year=$4, price=$5, total_page=$6, thickness=$7, category_id=$8, modified_at=$9, modified_by=$10 WHERE id=$11`,
		b.Title, b.Description, b.ImageURL, b.ReleaseYear, b.Price, b.TotalPage, b.Thickness, b.CategoryID, time.Now(), "admin", id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Update failed"})
		return
	}

	count, _ := res.RowsAffected()
	if count == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Updated successfull"})
}
