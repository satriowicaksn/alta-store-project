package controllers

import (
	"alta-store-project/lib/database"
	"alta-store-project/middlewares"
	"alta-store-project/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAllVoucherController(c echo.Context) error {
	vouchers, err := database.GetAllVoucher()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, models.Response{
		Status:  "success",
		Message: "success get all vouchers",
		Data:    vouchers,
	})
}

func GetMyVoucherController(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(c)
	vouchers, err := database.GetMyVoucher(userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	} else if vouchers == false {
		return c.JSON(http.StatusOK, models.Response{
			Status:  "fail",
			Message: "You don't have a voucher, make a purchase to get it",
			Data:    vouchers,
		})
	}
	return c.JSON(http.StatusOK, models.Response{
		Status:  "success",
		Message: "success get your vouchers",
		Data:    vouchers,
	})
}
