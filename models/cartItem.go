package models

import "time"

type Cart_item struct {
	Cart_item_id uint64    `json: "cart_item_id" gorm: "primaryKey"`
	Cart_id      int       `json: "cart_id" form: "cart_id"`
	Product_id   int       `json: "product_id" form: "product_id"`
	Qty          int       `json: "qty" form:"qty"`
	Price        int       `json: "price" form:"price"`
	Created_at   time.Time `json: "created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
