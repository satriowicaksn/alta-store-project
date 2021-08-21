package database

import (
	"alta-store-project/config"
	"alta-store-project/models"
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
