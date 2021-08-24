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
