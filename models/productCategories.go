package models

import "time"

type Product_categories struct {
	Category_id   int    `json:"category_id" form:"category_id"`
	Category_name string `json:"category_name" form:"category_name"`
	Created_at    time.Time
	Updated_at    time.Time
}
