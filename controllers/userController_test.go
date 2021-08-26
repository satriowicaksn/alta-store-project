package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserDetailController(t *testing.T) {
	t.Run("test case 1, valid user id", func(t *testing.T) {
		e := InitEcho()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath("/users/1")

		// Assertions
		if assert.NoError(t, GetUserDetailController(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			// body := rec.Body.String()
			// assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartsWith))
		}
	})
	t.Run("test case 2, invalid user id", func(t *testing.T) {
		e := InitEcho()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/users/18")

		// Assertions
		if assert.NoError(t, GetUserDetailController(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			// body := rec.Body.String()
			// assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartsWith))
		}
	})
}
