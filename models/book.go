package models

import (
	"time"
)

type Book struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	CategoryID  int       `json:"category_id"`
	Description string    `json:"description"`
	ImageURL    string    `json:"image_url"`
	ReleaseYear int       `json:"release_year"`
	Price       float64   `json:"price"`
	TotalPage   int       `json:"total_page"`
	Thickness   float64   `json:"thickness"`
	CreatedAt   time.Time `json:"created_at"`
	CreatedBy   int       `json:"created_by"`
	ModifiedAt  time.Time `json:"modified_at"`
	ModifiedBy  time.Time `json:"modified_by"`
}
