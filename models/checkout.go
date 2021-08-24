package models

type Checkout struct {
	Product    string `json:"product"`
	Item_total int    `json:"item_total"`
	Amount     int    `json:"amount"`
}
