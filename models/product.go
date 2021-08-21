package models

import "time"

type Product struct {
	Product_id          int    `json:"product_id" form:"product_id"`
	Product_category_id int    `json:"product_categories_id" form:"product_categories_id"`
	Product_name        string `json:"product_name" form:"product_name"`
	Description         string `json:"description" form:"description"`
	Price               int    `json:"price" form:"price"`
	Stock               int    `json:"stock" form:"stock"`
	Product_status      int    `json:"product_status" form:"product_status"`
	Created_at          time.Time
	Updated_at          time.Time
}
