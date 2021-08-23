package models

import "time"

type Cart_item struct {
	Cart_item_id uint `gorm: "primaryKey"`
	Cart_id      int  `json: "cart_id" form: "cart_id"`
	Product_id   int  `json: "product_id" form: "product_id"`
	Qty          int  `json: "qty" form:"qty"`
	Price        int  `json: "price" form:"price"`
	Created_at   time.Time
	Updated_at   time.Time
}
