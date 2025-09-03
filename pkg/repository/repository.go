package repository

import (
	"github.com/jmoiron/sqlx"
	"storeApi/models"
)

type Store interface {
	CreateProduct(product models.Product) (bool, error)
	AddCountProduct(productId int, count int) (int, error)
	GetProducts() ([]models.Product, error)
	GetProductById(productId int) (models.Product, error)
	GetProductsByIds(productIds []int) ([]models.Product, error)
	DeleteProductById(productId int) (models.Product, error)
	UpdateProductById(productId int, product models.Product) (bool, error)
	ReduceCountProduct(productId int, count int) (models.Product, float64, error)
	CreateOrder(order models.OrderRequest) (int, error)
	GetOrderById(orderId int) (models.Order, error)
}

type Repository struct {
	Store
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Store: NewStoreMySql(db),
	}
}
