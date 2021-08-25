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
	// var testCases = []struct {
	// 	name                 string
	// 	path                 string
	// 	expectStatus         int
	// 	expectBodyStartsWith string
	// }{
	// 	{
	// 		name:                 "berhasil",
	// 		path:                 "/products",
	// 		expectBodyStartsWith: "{\"products\":[",
	// 		expectStatus:         http.StatusOK,
	// 	},
	// }

	// e := InitEcho()
	// req := httptest.NewRequest(http.MethodGet, "/", nil)
	// rec := httptest.NewRecorder()
	// c := e.NewContext(req, rec)

	// for _, testCase := range testCases {
	// 	c.SetPath(testCase.path)

	// 	// Assertions
	// 	if assert.NoError(t, GetProductControllers(c)) {
	// 		assert.Equal(t, http.StatusOK, rec.Code)
	// 		body := rec.Body.String()
	// 		assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartsWith))
	// 	}
	// }

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
