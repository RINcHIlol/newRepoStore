package models

import "database/sql"

type Product struct {
	ID          int             `json:"id,omitempty" db:"id"`
	Name        string          `json:"name" binding:"required" db:"name"`
	Price       float64         `json:"price" binding:"required" db:"price"`
	Description string          `json:"description" binding:"required" db:"description"`
	Image       []byte          `json:"image" binding:"required" db:"image"`
	Count       sql.NullInt64   `json:"count" db:"count"`
}
