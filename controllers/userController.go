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
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "login successfull",
		"token":  user.Token,
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
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "fail",
			"message": field + " field is required !!",
		})
	}

	// validasi registered email
	validateEmail, err := database.ValidateEmail(u.Email)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	} else if !validateEmail {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "fail",
			"message": "Email is already registered, please use another email",
		})
	}

	// validasi registered phone number
	validatePhone, err := database.ValidatePhone(u.Phone)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	} else if !validatePhone {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "fail",
			"message": "Phone number is already registered, please use another phone number",
		})
	}

	// input ke database
	registered, err := database.RegisterUser(u)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"status": "Register success",
		"data":   registered,
	})
}

func GetUserDetailController(c echo.Context) error {

	// penambahan
	params := c.Param("id")
	id := 0
	if params != "" {
		idInt, e := strconv.Atoi(params)
		if e != nil {
			return echo.NewHTTPError(http.StatusBadRequest, e.Error())
		}
		id = idInt
	}
	// end penambahan

	users, err := database.GetDetailUsers((id))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if users == false {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "fail",
			"message": "user with requested ID was not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  users,
	})
}
