package service

import (
	"fmt"
	"gopkg.in/mail.v2"
	"os"
	"storeApi"
	"storeApi/models"
	"storeApi/pkg/mailer"
	"storeApi/pkg/repository"
)

type StoreService struct {
	repo repository.Store
}

func NewStoreService(repo repository.Store) *StoreService {
	return &StoreService{repo: repo}
}

func (s *StoreService) AddNewProduct(product models.Product) (bool, error) {
	return s.repo.CreateProduct(product)
}

func (s *StoreService) AddCountProduct(productId int, count int) (int, error) {
	return s.repo.AddCountProduct(productId, count)
}

func (s *StoreService) GetProducts() ([]models.Product, error) {
	return s.repo.GetProducts()
}

func (s *StoreService) GetProductById(productId int) (models.Product, error) {
	return s.repo.GetProductById(productId)
}

func (s *StoreService) BuyProduct(orderReq models.OrderRequest) (bool, error) {
	var finalPrice float64
	orderId, err := s.repo.CreateOrder(orderReq)
	if err != nil {
		return false, err
	}

	order, err := s.repo.GetOrderById(orderId)

	var products []models.ProductCount
	for i := 0; i < len(orderReq.Products); i++ {
		//product, err := s.repo.DeleteProductById(orderReq.IdsProduct[i])
		product, price, err := s.repo.ReduceCountProduct(orderReq.Products[i].ID, orderReq.Products[i].Count)
		if err != nil {
			return false, err
		}
		finalPrice += price
		products = append(products, models.ProductCount{product, orderReq.Products[i].Count, price})
	}

	err = storeApi.PaymentEvent(int(finalPrice * 100)) // сумма в центах
	if err != nil {
		return false, fmt.Errorf("оплата не прошла: %v", err)
	}

	customerMsg, customerFiles, err := mailer.MailToCustomer(products, order, finalPrice)
	if err != nil {
		return false, err
	}

	sellerMsg, sellerFiles, err := mailer.MailToSeller(products, order, finalPrice)
	if err != nil {
		return false, err
	}

	d := mail.NewDialer("smtp.gmail.com", 587, "galimatron229@gmail.com", "wplgvcwvcsvxxfxp")
	if err := d.DialAndSend(customerMsg, sellerMsg); err != nil {
		return false, err
	}
	for _, f := range customerFiles {
		f.Close()
		os.Remove(f.Name())
	}
	for _, f := range sellerFiles {
		f.Close()
		os.Remove(f.Name())
	}

	return true, nil
}

func (s *StoreService) UpdateProductById(productId int, product models.Product) (bool, error) {
	return s.repo.UpdateProductById(productId, product)
}
