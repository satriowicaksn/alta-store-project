package database

import (
	"alta-store-project/config"
	"alta-store-project/models"
)

func GetCarts() (interface{}, error) {
	var cart []models.Cart

	if err := config.DB.Find(&cart).Error; err != nil {
		return nil, err
	}
	return cart, nil
}
