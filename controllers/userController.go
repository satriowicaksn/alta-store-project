package controllers

import (
	"alta-store-project/lib/database"
	"alta-store-project/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func LoginUserController(c echo.Context) error {
	user := models.Users{}
	c.Bind(&user)
	_, e := database.LoginUsers(&user)
	if e != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, e.Error())
	}
	return c.JSON(http.StatusOK, models.Response{
		Status:  "success",
		Message: "success login to your account",
		Data:    user.Token,
	})
}

func RegisterUserController(c echo.Context) error {
	u := new(models.Users)
	if err := c.Bind(u); err != nil {
		return err
	}

	// Validasi input harus terisi semuanya
	validateInput, field := database.ValidateInput(u)
	if !validateInput {
		return c.JSON(http.StatusBadRequest, models.Response{
			Status:  "fail",
			Message: field + " field is required !!",
		})
	}

	// validasi registered email
	validateEmail, err := database.ValidateEmail(u.Email)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	} else if !validateEmail {
		return c.JSON(http.StatusBadRequest, models.Response{
			Status:  "fail",
			Message: "Email is already registered, please use another email",
		})
	}

	// validasi registered phone number
	validatePhone, err := database.ValidatePhone(u.Phone)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	} else if !validatePhone {
		return c.JSON(http.StatusBadRequest, models.Response{
			Status:  "fail",
			Message: "Phone number is already registered, please use another phone number",
		})
	}

	// input ke database
	registered, err := database.RegisterUser(u)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, models.Response{
		Status:  "success",
		Message: "success to registered your account",
		Data:    registered,
	})
}

func GetUserDetailController(c echo.Context) error {

	// penambahan
	params := c.Param("id")
	id := 0
	if params != "" {
		idInt, e := strconv.Atoi(params)
		if e != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, e.Error())
		}
		id = idInt
	}
	// end penambahan

	users, err := database.GetDetailUsers((id))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if users == false {
		return c.JSON(http.StatusBadRequest, models.Response{
			Status:  "fail",
			Message: "user with requested ID was not found",
		})
	}

	return c.JSON(http.StatusOK, models.Response{
		Status:  "success",
		Message: "success getting data",
		Data:    users,
	})
}
