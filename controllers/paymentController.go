package controllers

import (
	"alta-store-project/lib/database"
	// "alta-store-project/middlewares"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetPaymentMethodController(c echo.Context) error {
	paymentMethod, err := database.GetPaymentMethod()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   paymentMethod,
	})
}
