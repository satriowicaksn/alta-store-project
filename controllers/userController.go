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
		"status":  "login successfull",
		"user_id": user.User_id,
		"token":   user.Token,
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
