package models

type Cart struct {
	Cart_id    uint `gorm:"primaryKey"`
	User_id    int  `json:"user_id" form:"user_id"`
	Cart_total int  `json:"cart_total" form:"cart_total"`
}