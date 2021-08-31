package controllers

import (
	"alta-store-project/lib/database"
	"alta-store-project/middlewares"
	"alta-store-project/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetCheckoutTotalController(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(c)
	checkouts, err := database.GetCheckoutTotal(userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if checkouts == false {
		return c.JSON(http.StatusOK, models.Response{
			Status:  "success",
			Message: "your shopping cart is empty",
		})
	}
	return c.JSON(http.StatusOK, models.Response{
		Status:  "success",
		Message: "success get checkout data from your shopping cart",
		Data:    checkouts,
	})
}

func GetCheckoutByIdController(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(c)
	cartItemId, _ := strconv.Atoi(c.Param("id"))
	checkout, err := database.GetCheckoutTotalById(cartItemId, userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if checkout == false {
		return c.JSON(http.StatusBadRequest, models.Response{
			Status:  "fail",
			Message: "your requested data was not found",
		})
	}
	return c.JSON(http.StatusOK, models.Response{
		Status:  "success",
		Message: "success get checkout data from your shopping cart",
		Data:    checkout,
	})
}

func PostCheckoutController(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(c)
	cartItemId := c.FormValue("cart_id")

	voucherCode := c.FormValue("voucher_code")

	// validasi jika menggunakan kode voucher
	if voucherCode != "" {
		validateVoucher, errString := database.ValidateUserVoucher(userId, voucherCode)
		if !validateVoucher {
			return c.JSON(http.StatusBadRequest, models.Response{
				Status:  "fail",
				Message: errString,
			})
		}
	}

	if cartItemId == "all" || cartItemId == "" {
		cartItemId = "0"
	}
	cartItemIdInt, _ := strconv.Atoi(cartItemId)
	checkout, err := database.CheckoutItem(cartItemIdInt, userId, voucherCode)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if checkout == false {
		return c.JSON(http.StatusBadRequest, models.Response{
			Status:  "fail",
			Message: "your requested data was not found",
		})
	}
	return c.JSON(http.StatusOK, models.Response{
		Status:  "success",
		Message: "success checkout product from your shopping cart",
		Data:    checkout,
	})
}
