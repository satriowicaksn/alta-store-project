package controllers

import (
	"alta-store-project/lib/database"
	"alta-store-project/middlewares"
	"net/http"
	"strconv"

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
		"data":   carts,
	})
}

func AddCartController(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(c)
	payloadData := make(map[string]string)
	payloadData["qty"] = c.FormValue("qty")
	payloadData["product_id"] = c.FormValue("product_id")

	carts, err := database.AddCartItems(payloadData, userId)
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
		"data":   carts,
	})
}

func DeleteCartController(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(c)
	cartItemId, _ := strconv.Atoi(c.Param("id"))
	validate, err := database.ValidateCartItems(cartItemId, userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if validate == false {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"status":  "fail",
			"message": "You do not have access to delete this data",
		})
	}
	database.DeleteCartItems(cartItemId)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "Product has been removed from your shopping cart",
	})
}
