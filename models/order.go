package models

import "time"

type Order struct {
	ID            int       `json:"id,omitempty" db:"id"`
	CustomerEmail string    `json:"customer_email" db:"customer_email"`
	Address       string    `json:"address" db:"address"`
	Price         float64   `json:"price" db:"price"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
}
