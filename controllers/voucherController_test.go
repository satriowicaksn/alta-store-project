package controllers

import (
	"net/http"
	"net/http/httptest"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllVoucherController(t *testing.T) {
	t.Run("test case 1, valid", func(t *testing.T) {
		e := InitEcho()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath("/vouchers")

		// Assertion
		if assert.NoError(t, GetAllVoucherController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})
}
