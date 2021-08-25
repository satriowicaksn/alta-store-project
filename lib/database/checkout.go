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

	cartItem := []models.Cart_item{}
	config.DB.Where("cart_id = ?", cartId).Find(&cartItem)
	checkout.Checkout_id = cartId
	checkout.Checkout_item = cartItem

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
	query := config.DB.Raw("SELECT cart_id AS checkout_id, qty AS item_total, (cart_items.price*qty) AS amount, product_name AS product FROM cart_items LEFT JOIN products ON cart_items.product_id = products.product_id WHERE cart_item_id = ?", cartItemId).Find(&checkout)

	cartItem := []models.Cart_item{}
	config.DB.Where("cart_item_id = ?", cartItemId).Find(&cartItem)
	checkout.Checkout_item = cartItem

	if query.Error != nil {
		return nil, query.Error
	}
	if query.RowsAffected == 0 {
		return false, nil
	}
	return checkout, nil
}

func CheckoutItem(cartItemId, userId int) (interface{}, error) {
	paymentId := CreatePaymentId()
	cartId := CheckCart(userId)

	// validasi apakah item yang direquest tersedia
	cartItem := []models.Cart_item{}
	// checkout all
	cekItem := config.DB.Where("cart_id = ?", cartId).Find(&cartItem)
	if cekItem.RowsAffected == 0 {
		return false, nil
	}

	// checkout by cart item id
	if cartItemId != 0 {
		cekItem := config.DB.Where("cart_item_id = ?", cartItemId).Find(&cartItem)
		if cekItem.RowsAffected == 0 {
			return false, nil
		}
	}

	// dapatkan data item total dan amount untuk dimasukkan ke informasi payment
	checkout := models.Checkout{}
	if cartItemId != 0 {
		// checkout by cart item id
		config.DB.Raw("SELECT qty AS item_total, (cart_items.price*qty) AS amount, product_name AS product FROM cart_items LEFT JOIN products ON cart_items.product_id = products.product_id WHERE cart_item_id = ?", cartItemId).Find(&checkout)
	} else {
		// checkout all
		config.DB.Raw("SELECT qty AS item_total, (cart_items.price*qty) AS amount, product_name AS product FROM cart_items LEFT JOIN products ON cart_items.product_id = products.product_id WHERE cart_id = ?", cartId).Find(&checkout)
	}

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

	// masukan checkout item ke tabel payment items dengan perulangan
	for _, item := range cartItem {
		paymentItem := models.Payment_item{
			Payment_id: paymentId,
			Product_id: item.Product_id,
			Price:      item.Price,
			Qty:        item.Qty,
		}
		payment.Payment_item = append(payment.Payment_item, paymentItem) // append ke slice agar tampil di response
		config.DB.Create(&paymentItem)
	}
	// end perulangan untuk input ke tabel payment items

	// hapus semua data dari cart item

	if cartItemId != 0 {
		// jika checkout all
		if err := config.DB.Where("cart_item_id = ?", cartItemId).Delete(&cartItem).Error; err != nil {
			return nil, err
		}
	} else {
		// jika checkout berdasarkan cart item id
		if err := config.DB.Where("cart_id = ?", cartId).Delete(&cartItem).Error; err != nil {
			return nil, err
		}
	}

	return payment, nil
}

// func CheckoutItemById(cartItemId, userId int) (interface{}, error) {
// 	paymentId := CreatePaymentId()

// 	// validasi apakah item yang direquest tersedia
// 	cartItem := models.Cart_item{}
// 	cekItem := config.DB.Where("cart_item_id = ?", cartItemId).Find(&cartItem)
// 	if cekItem.RowsAffected == 0 {
// 		return false, nil
// 	}

// 	// dapatkan data item total dan amount untuk dimasukkan ke informasi payment
// 	checkout := models.Checkout{}
// 	config.DB.Raw("SELECT qty AS item_total, (cart_items.price*qty) AS amount, product_name AS product FROM cart_items LEFT JOIN products ON cart_items.product_id = products.product_id WHERE cart_item_id = ?", cartItemId).Find(&checkout)

// 	// masukkan informasi payment ke tabel payments
// 	payment := models.Payment{
// 		Payment_id:     paymentId,
// 		User_id:        userId,
// 		Amount:         checkout.Amount,
// 		Payment_status: 0,
// 		Created_at:     time.Now(),
// 		Expired_at:     time.Now().AddDate(0, 0, 1),
// 	}
// 	config.DB.Create(&payment)

// 	// masukkan checkout item ke tabel payment items
// 	paymentItem := models.Payment_item{
// 		Payment_id:   paymentId,
// 		Product_id:   cartItem.Product_id,
// 		Product_name: checkout.Product,
// 		Price:        cartItem.Price,
// 		Qty:          cartItem.Qty,
// 	}
// 	payment.Payment_item = append(payment.Payment_item, paymentItem)
// 	config.DB.Create(&paymentItem)

// 	// kosongkan cart item sesuai id yang di checkout
// 	if err := config.DB.Where("cart_item_id = ?", cartItemId).Delete(&cartItem).Error; err != nil {
// 		return nil, err
// 	}
// 	return payment, nil
// }
