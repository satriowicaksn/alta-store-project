package controllers

import (
	"alta-store-project/lib/database"
	"alta-store-project/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetProductControllers(c echo.Context) error {
	products, err := database.GetProducts()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, models.Response{
		Status:  "success",
		Message: "success get all products",
		Data:    products,
	})
}

func GetProductsByCategoryControllers(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	products, err := database.GetProductsByCategory(id)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if products == false {
		return c.JSON(http.StatusBadRequest, models.Response{
			Status:  "fail",
			Message: "products with requested category was not found",
		})
	}
	return c.JSON(http.StatusOK, models.Response{
		Status:  "success",
		Message: "success get products by category id",
		Data:    products,
	})
}
