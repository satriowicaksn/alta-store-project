package database

import (
	"alta-store-project/config"
	"alta-store-project/models"
)

func GetCheckoutTotal(userId int) (interface{}, error) {
	checkout := models.Checkout{}
	cartId := CheckCart(userId)
	query := config.DB.Raw("SELECT SUM(qty) AS item_total, SUM(price*qty) AS amount FROM cart_items WHERE cart_id = ?", cartId).Find(&checkout)
	checkout.Product = "All in your cart"
	if query.Error != nil {
		return nil, query.Error
	}
	if query.RowsAffected == 0 {
		return false, nil
	}
	return checkout, nil
}

func GetCheckoutTotalById(cartItemId, userId int) (interface{}, error) {
	checkout := models.Checkout{}
	query := config.DB.Raw("SELECT qty AS item_total, (cart_items.price*qty) AS amount, product_name AS product FROM cart_items LEFT JOIN products ON cart_items.product_id = products.product_id WHERE cart_item_id = ?", cartItemId).Find(&checkout)
	if query.Error != nil {
		return nil, query.Error
	}
	if query.RowsAffected == 0 {
		return false, nil
	}
	return checkout, nil
}
