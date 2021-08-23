package models

import "time"

type Cart_item struct {
	ID         uint64    `json: "id" sql: "AUTO_INCREMENT" gorm: "primary_key,column:cart_item_id"`
	Cart_id    int       `json: "cart_id" form: "cart_id"`
	Product_id int       `json: "product_id" form: "product_id"`
	Qty        int       `json: "qty" form:"qty"`
	Price      int       `json: "price" form:"price"`
	Created_at time.Time `json: "created_at" gorm:"column:created_at" sql:"DEFAULT:current_timestamp"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"column:updated_at" sql:"DEFAULT:current_timestamp"`
}
