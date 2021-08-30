package controllers

import (
	"alta-store-project/lib/database"
	"alta-store-project/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetProductCategoriesControllers(c echo.Context) error {
	categories, err := database.GetProductCategories()

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, models.Response{
		Status:  "success",
		Message: "success get product categories",
		Data:    categories,
	})
}

func GetProductCategoriesByIdControllers(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	categories, err := database.GetProductCategoriesById(id)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if categories == false {
		return c.JSON(http.StatusBadRequest, models.Response{
			Status:  "fail",
			Message: "requested category was not found",
		})
	}
	return c.JSON(http.StatusOK, models.Response{
		Status:  "success",
		Message: "success getting data",
		Data:    categories,
	})
}

func PostProductCategories(c echo.Context) error {
	payloadData := make(map[string]string)
	payloadData["category_name"] = c.FormValue("category_name")
	categories, err := database.CreateProductCategories(payloadData)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"status":  "success",
		"message": "new category has created",
		"data":    categories,
	})
}

func PutProductCategoriesById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	payloadData := make(map[string]string)
	payloadData["category_name"] = c.FormValue("category_name")
	categories, err := database.UpdateProductCategories(payloadData, id)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "category has updated",
		"data":    categories,
	})
}

func DeleteProductCategoriesById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	database.DeleteProductCategories(id)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "category has deleted",
	})
}
