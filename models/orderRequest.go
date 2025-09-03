package models

type OrderRequest struct {
	CustomerEmail string           `json:"customer_email"`
	Address       string           `json:"address"`
	Products      []ProductRequest `json:"products"`
}

type ProductRequest struct {
	ID    int `json:"id_product"`
	Count int `json:"count"`
}
