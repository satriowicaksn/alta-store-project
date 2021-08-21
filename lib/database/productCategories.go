package database

import (
	"alta-store-project/config"
	"alta-store-project/models"
	"time"
)

func GetProductCategories() (interface{}, error) {
	var productCategories []models.Product_categories

	if e := config.DB.Find(&productCategories).Error; e != nil {
		return nil, e
	}
	return productCategories, nil

}

func GetProductCategoriesById(id int) (interface{}, error) {
	var productCategories []models.Product_categories
	query := config.DB.Where("category_id = ?", id).Find(&productCategories)
	if e := query.Error; e != nil {
		return nil, e
	}
	if query.RowsAffected == 0 {
		return false, nil
	}
	return productCategories, nil
}

func CreateProductCategories(payloadData map[string]string) (interface{}, error) {
	category := models.Product_categories{Category_name: payloadData["category_name"], Created_at: time.Now()}
	if err := config.DB.Create(&category).Error; err != nil {
		return nil, err
	}
	return category, nil
}

func UpdateProductCategories(payloadData map[string]string, id int) (interface{}, error) {
	category := models.Product_categories{Category_name: payloadData["category_name"], Updated_at: time.Now()}
	if err := config.DB.Where("category_id = ?", id).Updates(&category).Error; err != nil {
		return nil, err
	}
	return category, nil
}

func DeleteProductCategories(id int) (interface{}, error) {
	var productCategories []models.Product_categories
	if err := config.DB.Where("category_id = ?", id).Delete(&productCategories).Error; err != nil {
		return nil, err
	}
	return true, nil
}
