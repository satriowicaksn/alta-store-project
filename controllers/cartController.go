package controllers

import (
	"alta-store-project/lib/database"
	"alta-store-project/middlewares"
	"alta-store-project/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetCartController(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(c)
	carts, err := database.GetCarts(userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if carts == false {
		return c.JSON(http.StatusOK, models.Response{
			Status:  "success",
			Message: "your shopping cart is empty",
		})
	}
	return c.JSON(http.StatusOK, models.Response{
		Status:  "success",
		Message: "success get your shopping cart",
		Data:    carts,
	})
}

func AddCartController(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(c)
	payloadData := make(map[string]string)
	payloadData["qty"] = c.FormValue("qty")
	payloadData["product_id"] = c.FormValue("product_id")

	carts, err := database.AddCartItems(payloadData, userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if carts == false {
		return c.JSON(http.StatusBadRequest, models.Response{
			Status:  "fail",
			Message: "product not found or out of stock",
		})
	}
	return c.JSON(http.StatusCreated, models.Response{
		Status:  "success",
		Message: "success add product to your shopping cart",
		Data:    carts,
	})
}

func DeleteCartController(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(c)
	cartItemId, _ := strconv.Atoi(c.Param("id"))
	validate, validateItem, err := database.ValidateCartItems(cartItemId, userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if validate == false {
		return c.JSON(http.StatusForbidden, models.Response{
			Status:  "fail",
			Message: "You do not have access to delete this data",
		})
	}
	database.ReturnStock(validateItem["product_id"], validateItem["qty"])
	database.DeleteCartItems(cartItemId)
	return c.JSON(http.StatusOK, models.Response{
		Status:  "success",
		Message: "Product has been removed from your shopping cart",
	})
}
