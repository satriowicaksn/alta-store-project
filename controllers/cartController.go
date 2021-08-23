package controllers

import (
	"alta-store-project/lib/database"
	"alta-store-project/middlewares"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetCartController(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(c)
	_, getUserErr := database.GetDetailUsers(userId)
	if getUserErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, getUserErr.Error())
	}
	carts, err := database.GetCarts(userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"cart":   carts,
		"ID":     userId,
	})

}
