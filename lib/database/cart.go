package database

import (
	"alta-store-project/config"
	"alta-store-project/models"
	"strconv"
)

func GetCarts(userId int) (interface{}, error) {
	var cart_item []models.Cart_item
	cartId := CheckCart(userId)
	query := config.DB.Where("cart_id = ?", cartId).Find(&cart_item)
	if err := query.Error; err != nil {
		return nil, err
	}
	if query.RowsAffected == 0 {
		return false, nil
	}
	return cart_item, nil
}

func CheckCart(userId int) int {
	cart := models.Cart{
		User_id:    userId,
		Cart_total: 0,
	}
	query := config.DB.Where("user_id = ?", userId).Find(&cart)
	if query.RowsAffected == 0 {
		config.DB.Create(&cart)
		CheckCart(userId)
	}
	return int(cart.Cart_id)
}

func AddCartItems(payloadData map[string]string, userId int) (interface{}, error) {

	cartId := CheckCart(userId)
	productId, _ := strconv.Atoi(payloadData["product_id"])
	productStock, productPrice, err := GetProductById(productId)
	qty, _ := strconv.Atoi(payloadData["qty"])
	if err != nil {
		return nil, err
	}

	// validasi cek apakah stock produk tidak ditemukan
	if productStock == 0 {
		return false, nil
	}

	// validasi cek apakah stock kurang
	if productStock < qty {
		return false, nil
	}

	// masukan ke cart
	cartItem := models.Cart_item{
		Cart_id:    cartId,
		Product_id: productId,
		Qty:        qty,
		Price:      productPrice,
	}
	addToCart := config.DB.Create(&cartItem)
	if addToCart.Error != nil {
		return nil, addToCart.Error
	}
	UpdateProductStockById(productId, productStock, qty)
	return cartItem, nil
}

func DeleteCartItems(cart_item_id int) (interface{}, error) {
	var cartItem []models.Cart_item
	if err := config.DB.Where("cart_item_id = ?", cart_item_id).Delete(&cartItem).Error; err != nil {
		return nil, err
	}
	return true, nil
}

func ValidateCartItems(cart_item_id, userId int) (bool, error) {
	cartId := CheckCart(userId)
	cartItem := models.Cart_item{}
	query := config.DB.Where("cart_item_id = ?", cart_item_id).Find(&cartItem)

	if query.Error != nil {
		return false, query.Error
	} else if cartItem.Cart_id != cartId || query.RowsAffected == 0 {
		return false, nil
	}
	return true, nil
}
