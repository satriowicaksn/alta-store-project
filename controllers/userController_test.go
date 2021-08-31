package controllers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"alta-store-project/config"
	"alta-store-project/models"

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

func TestRegisterUserController(t *testing.T) {
	t.Run("Test case 1, valid register", func(t *testing.T) {

		var user []models.Users
		test := "testing"
		config.DB.Where("email = ?", test).Delete(&user)
		e := InitEcho()
		userJSON := `{"name":"testing","email":"testing","phone":"012345","password":"testing","address":"malang"}`
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, RegisterUserController(c)) {
			assert.Equal(t, http.StatusCreated, rec.Code)
		}
	})
	t.Run("Test case 2, uncompleted email", func(t *testing.T) {

		var user []models.Users
		test := "testing"
		config.DB.Where("email = ?", test).Delete(&user)
		e := InitEcho()
		userJSON := `{"name":"testing","email":"","phone":"","password":"","address":"malang"}`
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, RegisterUserController(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})
	t.Run("Test case 2.1, uncompleted name", func(t *testing.T) {

		var user []models.Users
		test := "testing"
		config.DB.Where("email = ?", test).Delete(&user)
		e := InitEcho()
		userJSON := `{"name":"","email":"","phone":"","password":"","address":"malang"}`
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, RegisterUserController(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})

	t.Run("Test case 2.2, uncompleted phone", func(t *testing.T) {

		var user []models.Users
		test := "testing"
		config.DB.Where("email = ?", test).Delete(&user)
		e := InitEcho()
		userJSON := `{"name":"testing","email":"testing","phone":"","password":"","address":"malang"}`
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, RegisterUserController(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})

	t.Run("Test case 2.3, uncompleted password", func(t *testing.T) {

		var user []models.Users
		test := "testing"
		config.DB.Where("email = ?", test).Delete(&user)
		e := InitEcho()
		userJSON := `{"name":"testing","email":"testing","phone":"0","password":"","address":"malang"}`
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, RegisterUserController(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})

	t.Run("Test case 2.4, uncompleted address", func(t *testing.T) {

		var user []models.Users
		test := "testing"
		config.DB.Where("email = ?", test).Delete(&user)
		e := InitEcho()
		userJSON := `{"name":"testing","email":"testing","phone":"0","password":"asas","address":""}`
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, RegisterUserController(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})

	t.Run("Test case 3, registered email", func(t *testing.T) {

		e := InitEcho()
		userJSON := `{"name":"testing","email":"satrio@gmail.com","phone":"012121212","password":"testing","address":"malang"}`
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, RegisterUserController(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})

	t.Run("Test case 4, registered phone", func(t *testing.T) {
		var user []models.Users
		test := "testing"
		config.DB.Where("email = ?", test).Delete(&user)
		e := InitEcho()
		userJSON := `{"name":"testing","email":"testing","phone":"081334304990","password":"testing","address":"malang"}`
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, RegisterUserController(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})
}
