package database

import (
	"alta-store-project/config"
	"alta-store-project/models"
	"strconv"
	"time"
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

func GetPaymentHistory(userId int) (interface{}, error) {
	var payment []models.Payment
	if err := config.DB.Where("user_id = ? AND payment_status = 1", userId).Find(&payment).Error; err != nil {
		return nil, err
	}
	return payment, nil
}

func GetPendingPayment(userId int) (interface{}, error) {
	var payment []models.Payment
	if err := config.DB.Where("user_id = ? AND payment_status = 0", userId).Find(&payment).Error; err != nil {
		return nil, err
	}
	return payment, nil
}

func PostPayment(userId int, payloadData map[string]string) (interface{}, error) {
	payment := models.Payment{}
	payment_id, _ := strconv.Atoi(payloadData["payment_id"])
	payment_method, _ := strconv.Atoi(payloadData["payment_method"])
	amount, _ := strconv.Atoi(payloadData["amount"])
	check := config.DB.Where("payment_id = ? AND amount = ? AND user_id = ?", payment_id, amount, userId).Find(&payment)
	if check.Error != nil {
		return nil, check.Error
	}
	if check.RowsAffected == 0 {
		return false, nil
	}
	paymentUpdate := models.Payment{
		Payment_method: payment_method,
		Payed_at:       time.Now(),
		Payment_status: 1,
	}
	config.DB.Where("payment_id = ?", payloadData["payment_id"]).Updates(&paymentUpdate)
	return paymentUpdate, nil
}
