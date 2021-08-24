package database

import (
	"alta-store-project/config"
	"alta-store-project/models"
	"time"
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

// func CheckoutAllItem(userId int) (interface{}, error) {
// 	paymentId := CreatePaymentId()
// 	cartId := CheckCart(userId)

// }

func CheckoutItemById(cartItemId, userId int) (interface{}, error) {
	paymentId := CreatePaymentId()
	cartItem := models.Cart_item{}
	cekItem := config.DB.Where("cart_item_id = ?", cartItemId).Find(&cartItem)

	// jika dicek tidak ditemukan ID nya maka return false
	if cekItem.RowsAffected == 0 {
		return false, nil
	}

	checkout := models.Checkout{}
	config.DB.Raw("SELECT qty AS item_total, (cart_items.price*qty) AS amount, product_name AS product FROM cart_items LEFT JOIN products ON cart_items.product_id = products.product_id WHERE cart_item_id = ?", cartItemId).Find(&checkout)

	// masukkan ke checkout item (tabel payment item)
	checkoutItem := models.Payment_item{
		Payment_id:   paymentId,
		Product_id:   cartItem.Product_id,
		Product_name: checkout.Product,
		Price:        cartItem.Price,
		Qty:          cartItem.Qty,
	}
	config.DB.Create(&checkoutItem)

	// masukkan informasi payment
	payment := models.Payment{
		Payment_id:     paymentId,
		User_id:        userId,
		Amount:         checkout.Amount,
		Payment_status: 0,
		Created_at:     time.Now(),
		Expired_at:     time.Now().AddDate(0, 0, 1),
	}
	config.DB.Create(&payment)

	// kosongkan cart item
	if err := config.DB.Where("cart_item_id = ?", cartItemId).Delete(&cartItem).Error; err != nil {
		return nil, err
	}
	return payment, nil
}
