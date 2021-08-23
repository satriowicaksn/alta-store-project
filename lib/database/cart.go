package database

import (
	"alta-store-project/config"
	"alta-store-project/models"
)

func GetCarts(userId int) (interface{}, error) {
	var cart []models.Cart

	if err := config.DB.Where("user_id = ?", userId).Find(&cart).Error; err != nil {
		return nil, err
	}
	return cart, nil
}
