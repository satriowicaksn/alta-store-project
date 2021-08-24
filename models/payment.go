package models

import "time"

type Payment struct {
	Payment_id     string    `json:"payment_id" gorm:"primaryKey"`
	Payment_method int       `json:"payment_method"`
	User_id        int       `json:"user_id"`
	Amount         int       `json:"amount"`
	Payment_status int       `json:"payment_status"`
	Created_at     time.Time `json:"created_at"`
	Payed_at       time.Time `json:"payed_at"`
	Expired_at     time.Time `json:"expired_at"`
}

type Uuid struct {
	Uuid string `json:"uuid"`
}
