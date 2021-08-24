package controllers

import (
	"alta-store-project/lib/database"
	"alta-store-project/middlewares"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetCheckoutTotalController(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(c)
	checkouts, err := database.GetCheckoutTotal(userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if checkouts == false {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status":  "success",
			"message": "your shopping cart is empty",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   checkouts,
	})
}

func GetCheckoutByIdController(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(c)
	cartItemId, _ := strconv.Atoi(c.Param("id"))
	checkout, err := database.GetCheckoutTotalById(cartItemId, userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if checkout == false {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "fail",
			"message": "your requested data was not found",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   checkout,
	})
}

func PostCheckoutByIdController(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(c)
	cartItemId, _ := strconv.Atoi(c.Param("id"))
	checkout, err := database.CheckoutItemById(cartItemId, userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if checkout == false {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "fail",
			"message": "your requested data was not found",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   checkout,
	})
}
