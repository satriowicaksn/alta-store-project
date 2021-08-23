package controllers

import (
	"alta-store-project/lib/database"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetCartController(c echo.Context) error {
	carts, err := database.GetCarts()
	userId := "hai"
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"cart":   carts,
		"ID":     userId,
	})

}
