package models

import "time"

type Category struct {
    ID         int       `json:"id"`
    Name       string    `json:"name"`
    CreatedAt  time.Time `json:"created_at"`
    CreatedBy  int       `json:"created_by"`
    ModifiedAt time.Time `json:"modified_at"`
    ModifiedBy int       `json:"modified_by"`
}
