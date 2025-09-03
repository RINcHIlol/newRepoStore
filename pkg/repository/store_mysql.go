package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"storeApi/models"
	"time"
)

type StoreMySql struct {
	db *sqlx.DB
}

func NewStoreMySql(db *sqlx.DB) *StoreMySql {
	return &StoreMySql{db: db}
}

func (p *StoreMySql) CreateProduct(product models.Product) (bool, error) {
	query := `
		INSERT INTO products (order_num, name, price, description, image, count)
		VALUES (?, ?, ?, ?, ?, ?)
	`

	count := 0
	if product.Count.Valid {
		count = int(product.Count.Int64)
	}
	_, err := p.db.Exec(query, product.Name, product.Price, product.Description, product.Image, count)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (p *StoreMySql) AddCountProduct(productId int, count int) (int, error) {
	query := `
		UPDATE products 
		SET count = count + ?
		WHERE id = ?
	`

	result, err := p.db.Exec(query, count, productId)
	if err != nil {
		return 0, err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return 0, err
	}

	product, err := p.GetProductById(productId)
	if err != nil {
		return 0, err
	}

	if product.Count.Valid {
		return int(product.Count.Int64), nil
	}
	return 0, nil
}

func (p *StoreMySql) GetProducts() ([]models.Product, error) {
	var products []models.Product

	query := `SELECT id, name, price, description, image, count FROM products`

	err := p.db.Select(&products, query)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (p *StoreMySql) GetProductById(productId int) (models.Product, error) {
	var product models.Product

	query := `SELECT id, name, price, description, image, count FROM products WHERE id = ?`

	err := p.db.Get(&product, query, productId)
	if err != nil {
		return product, err
	}

	return product, nil
}

func (p *StoreMySql) GetProductsByIds(productIds []int) ([]models.Product, error) {
	var result []models.Product
	for _, id := range productIds {
		product, err := p.GetProductById(id)
		if err != nil {
			return nil, err
		}
		result = append(result, product)
	}
	return result, nil
}

func (p *StoreMySql) DeleteProductById(productId int) (models.Product, error) {
	var product models.Product

	querySelect := `SELECT id, name, price, description, image, count FROM products WHERE id = ?`
	err := p.db.Get(&product, querySelect, productId)
	if err != nil {
		return models.Product{}, err
	}

	queryDelete := `DELETE FROM products WHERE id = ?`
	result, err := p.db.Exec(queryDelete, productId)
	if err != nil {
		return models.Product{}, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return models.Product{}, err
	}
	if rowsAffected == 0 {
		return models.Product{}, fmt.Errorf("no product with id %d found", productId)
	}

	return product, nil
}

func (p *StoreMySql) UpdateProductById(productId int, product models.Product) (bool, error) {
	query := `
		UPDATE products 
		SET name = ?, price = ?, description = ?, image = ?, count = ?
		WHERE id = ?
	`

	count := 0
	if product.Count.Valid {
		count = int(product.Count.Int64)
	}
	result, err := p.db.Exec(query, product.Name, product.Price, product.Description, product.Image, count, productId)
	if err != nil {
		return false, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	return rowsAffected > 0, nil
}

func (p *StoreMySql) ReduceCountProduct(productId int, count int) (models.Product, float64, error) {
	var product models.Product
	query := `
		UPDATE products 
		SET count = count - ?
		WHERE id = ?
	`

	result, err := p.db.Exec(query, count, productId)
	if err != nil {
		return product, 0, err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return product, 0, err
	}

	product, err = p.GetProductById(productId)
	if err != nil {
		return product, 0, err
	}

	price := product.Price * float64(count)

	return product, price, nil
}

func (p *StoreMySql) CreateOrder(order models.OrderRequest) (int, error) {
	var finalPrice float64
	var semiPrice float64
	for i := 0; i < len(order.Products); i++ {
		query := `SELECT price FROM products WHERE id = ?`

		err := p.db.Get(&semiPrice, query, order.Products[i].ID)
		if err != nil {
			return 0, err
		}
		finalPrice += (semiPrice * float64(order.Products[i].Count))
	}

	query := `
		INSERT INTO orders (customer_email, address, created_at, price)
		VALUES (?, ?, ?, ?)
	`

	result, err := p.db.Exec(query, order.CustomerEmail, order.Address, time.Now(), finalPrice)
	if err != nil {
		return 0, err
	}

	orderID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	for i := 0; i < len(order.Products); i++ {
		query := `
		INSERT INTO order_items (order_id, product_id)
		VALUES (?, ?)
	`

		_, err := p.db.Exec(query, orderID, order.Products[i].ID)
		if err != nil {
			return 0, err
		}
	}

	return int(orderID), nil
}

func (p *StoreMySql) GetOrderById(orderId int) (models.Order, error) {
	var order models.Order

	query := `SELECT id, customer_email, address, created_at FROM orders WHERE id = ?`

	err := p.db.Get(&order, query, orderId)
	if err != nil {
		return order, err
	}

	return order, nil
}
