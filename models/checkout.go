package models

type Checkout struct {
	Checkout_id   int         `json:"checkout_id"`
	Product       string      `json:"product"`
	Item_total    int         `json:"item_total"`
	Amount        int         `json:"amount"`
	Checkout_item []Cart_item `json:"items" gorm:"foreignKey:Cart_id;references:Checkout_id"`
}
