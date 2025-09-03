package models

type ProductCount struct {
	Product Product `json:"product"`
	Count   int     `json:"count"`
	Price   float64 `json:"price"`
}
