package controllers

import (
	"alta-store-project/config"
	"net/http"
	"net/http/httptest"

	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func InitEcho() *echo.Echo {
	// Setup
	config.InitDB()
	e := echo.New()
	return e
}

func TestGetProductControllers(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		e := InitEcho()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath("/products")

		// Assertions
		if assert.NoError(t, GetProductControllers(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			// body := rec.Body.String()
			// assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartsWith))
		}
	})
}
func TestGetProductsByCategoryControllers(t *testing.T) {
	t.Run("Test Case 1, Valid Category Id", func(t *testing.T) {
		e := InitEcho()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath("/products/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")

		// Assertions
		if assert.NoError(t, GetProductsByCategoryControllers(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})
	t.Run("Test Case 2, Invalid Category Id", func(t *testing.T) {
		e := InitEcho()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath("/products/:id")
		c.SetParamNames("id")
		c.SetParamValues("19")

		// Assertions
		if assert.NoError(t, GetProductsByCategoryControllers(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})
}
