package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func WelcomeControllers(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome to Alta Store")
}
