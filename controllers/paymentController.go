package controllers

import (
	"alta-store-project/lib/database"
	"alta-store-project/middlewares"
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

func PostPaymentController(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(c)
	payloadData := make(map[string]string)
	payloadData["payment_id"] = c.FormValue("payment_id")
	payloadData["amount"] = c.FormValue("amount")
	payloadData["payment_method"] = c.FormValue("payment_method")
	payment, err := database.PostPayment(userId, payloadData)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if payment == "Payment method invalid" || payment == "This bill has been paid" || payment == "The bill you want to pay was not found" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "fail",
			"message": payment,
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   payment,
	})
}

func GetPendingPaymentController(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(c)
	payment, err := database.GetPendingPayment(userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	} else if payment == false {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status":  "success",
			"message": "You don't have pending payment",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   payment,
	})
}

func GetPaymentHistoryController(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(c)
	payment, err := database.GetPaymentHistory(userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	} else if payment == false {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status":  "success",
			"message": "You don't have payment history",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   payment,
	})
}
