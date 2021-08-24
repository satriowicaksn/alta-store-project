package models

type Payment_item struct {
	Payment_item_id int    `json:"payment_item_id"`
	Payment_id      string `json:"payment_id"`
	Product_id      int    `json:"product_id"`
	Product_name    string `json:"product_name"`
	Price           int    `json:"price"`
	Qty             int    `json:"qty"`
}
