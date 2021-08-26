package controllers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetUserDetailController(t *testing.T) {
	t.Run("Test case 1, valid user id", func(t *testing.T) {
		e := InitEcho()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath("/users/:id")
		c.SetParamNames("id")
		c.SetParamValues("2")

		// Assertions
		if assert.NoError(t, GetUserDetailController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			// body := rec.Body.String()
			// assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartsWith))
		}
	})
	t.Run("Test case 2, invalid user id", func(t *testing.T) {
		e := InitEcho()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath("/users/:id")
		c.SetParamNames("id")
		c.SetParamValues("99")

		// Assertions
		if assert.NoError(t, GetUserDetailController(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			// body := rec.Body.String()
			// assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartsWith))
		}
	})
}

func TestLoginUserController(t *testing.T) {
	t.Run("Test case 1, valid authentication user", func(t *testing.T) {
		e := InitEcho()
		userJSON := `{"email":"satrio@gmail.com","password":"satrio12345"}`
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, LoginUserController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}

	})
}
