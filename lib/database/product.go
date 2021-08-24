package database

import (
	"alta-store-project/config"
	"alta-store-project/models"
	"time"
)

func GetProducts() (interface{}, error) {
	var products []models.Product

	if e := config.DB.Find(&products).Error; e != nil {
		return nil, e
	}
	return products, nil
}

func GetProductsByCategory(id int) (interface{}, error) {
	var products []models.Product
	query := config.DB.Where("product_category_id = ?", id).Find(&products)
	if e := query.Error; e != nil {
		return nil, e
	}
	if query.RowsAffected == 0 {
		return false, nil
	}
	return products, nil
}

func GetProductById(id int) (int, int, error) {
	products := models.Product{}
	query := config.DB.Where("product_id = ?", id).Find(&products)
	if e := query.Error; e != nil {
		return 0, 0, e
	}
	if query.RowsAffected == 0 {
		return 0, 0, nil
	}
	return products.Stock, products.Price, nil
}

func UpdateProductStockById(id, stock, qty int) {
	product := models.Product{
		Stock:      stock - qty,
		Updated_at: time.Now(),
	}
	config.DB.Where("product_id = ?", id).Updates(&product)
}

func ReturnStock(product_id, qty int) {
	product := models.Product{}
	stock := 0
	getQuery := config.DB.Where("product_id = ?", product_id).Find(&product)

	if getQuery.RowsAffected > 0 {
		stock = product.Stock
		newStock := models.Product{
			Stock:      stock + qty,
			Updated_at: time.Now(),
		}
		config.DB.Where("product_id = ?", product_id).Updates(&newStock)
	}
}
