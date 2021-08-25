package models

import "time"

type Payment struct {
	Payment_id     string         `json:"payment_id" gorm:"primaryKey"`
	Payment_method int            `json:"payment_method"`
	User_id        int            `json:"user_id"`
	Amount         int            `json:"amount"`
	Payment_status int            `json:"payment_status"`
	Created_at     time.Time      `json:"created_at"`
	Payed_at       time.Time      `json:"payed_at"`
	Expired_at     time.Time      `json:"expired_at"`
	Payment_item   []Payment_item `json:"payment_item" gorm:"foreignKey:Payment_id;references:Payment_id"`
}

type Uuid struct {
	Uuid string `json:"uuid"`
}

type Payment_method struct {
	Payment_method_id     int    `json:"payment_method_id" gorm:"primaryKey"`
	Payment_method_name   string `json:"payment_method_name"`
	Payment_method_number string `json:"payment_method_number"`
}
