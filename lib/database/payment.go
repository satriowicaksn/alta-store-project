package database

import (
	"alta-store-project/config"
	"alta-store-project/models"
)

func CreatePaymentId() string {
	UniqueID := models.Uuid{}
	config.DB.Raw("SELECT uuid() as uuid").Scan(&UniqueID)
	return UniqueID.Uuid
}

func GetPaymentMethod() (interface{}, error) {
	var paymentMethod []models.Payment_method
	if err := config.DB.Find(&paymentMethod).Error; err != nil {
		return nil, err
	}
	return paymentMethod, nil
}
