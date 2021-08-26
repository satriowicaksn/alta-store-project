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

type Payment_item struct {
	Payment_item_id int    `json:"payment_item_id" gorm:"primaryKey"`
	Payment_id      string `json:"payment_id"`
	Product_id      int    `json:"product_id"`
	Price           int    `json:"price"`
	Qty             int    `json:"qty"`
}

type Payment_history struct {
	Payment_id          string    `json:"payment_id"`
	Payment_method_name string    `json:"payment_method"`
	Amount              int       `json:"amount"`
	Payed_at            time.Time `json:"payed_at"`
}

type Uuid struct {
	Uuid string `json:"uuid"`
}

type Payment_method struct {
	Payment_method_id     int    `json:"payment_method_id" gorm:"primaryKey"`
	Payment_method_name   string `json:"payment_method_name"`
	Payment_method_number string `json:"payment_method_number"`
}
