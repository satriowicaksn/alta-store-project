package controllers

import (
	"alta-store-project/lib/database"
	"alta-store-project/middlewares"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetCartController(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(c)
	carts, err := database.GetCarts(userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if carts == false {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status":  "success",
			"message": "your shopping cart is empty",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"cart":   carts,
	})
}

func AddCartController(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(c)
	payloadData := make(map[string]string)
	payloadData["qty"] = c.FormValue("qty")
	payloadData["product_id"] = c.FormValue("product_id")

	carts, err := database.AddCartItems(payloadData, userId)
	// cek := 0
	// if cek == 0 {
	// 	return c.JSON(http.StatusOK, map[string]interface{}{
	// 		"status":  "fail",
	// 		"message": "product not found or out of stock",
	// 	})
	// }

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if carts == false {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status":  "fail",
			"message": "product not found or out of stock",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"cart":   carts,
	})
}
