package models

import "time"

type Voucher struct {
	Voucher_id          int       `json:"voucher_id" gorm:"primaryKey"`
	Voucher_code        string    `json:"voucher_code"`
	Description         string    `json:"description"`
	Voucher_type        int       `json:"voucher_type"`
	Voucher_disc        int       `json:"disc"`
	Minimum_transaction int       `json:"minimum_transaction"`
	Created_at          time.Time `json:"created_at"`
	Updated_at          time.Time `json:"updated_at"`
}

type User_voucher struct {
	User_voucher_id int `json:"user_voucher_id" gorm:"primaryKey"`
	User_id         int `json:"user_id"`
	Voucher_id      int `json:"voucher_id"`
	Status          int `json:"status"`
}

type Payment_voucher struct {
	Payment_id   string `json:"payment_id"`
	Voucher_code string `json:"voucher_code"`
	Total_bill   int    `json:"total_bill"`
	Disc         int    `json:"discount"`
	Final_bill   int    `json:"final_bill"`
}

type My_voucher struct {
	Voucher_code        string `json:"voucher_code"`
	Description         string `json:"description"`
	Voucher_disc        int    `json:"disc"`
	Minimum_transaction int    `json:"minimum_transaction"`
}
