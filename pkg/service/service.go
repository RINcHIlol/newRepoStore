package service

import (
	"storeApi/models"
	"storeApi/pkg/repository"
)

type Store interface {
	AddNewProduct(product models.Product) (bool, error)
	AddCountProduct(productId int, count int) (int, error)
	GetProducts() ([]models.Product, error)
	GetProductById(productId int) (models.Product, error)
	BuyProduct(order models.OrderRequest) (int, error)
	UpdateProductById(productId int, product models.Product) (bool, error)
	DeleteProductById(productId int) (models.Product, error)
}

type Service struct {
	Store
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Store: NewStoreService(repos.Store),
	}
}
