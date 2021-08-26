package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCartController(t *testing.T) {
	t.Run("Test case 1, valid test", func(t *testing.T) {
		e := InitEcho()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		var bearer = "Bearer " + "ansadhahfbhahfbahfba"
		req.Header.Add("Authorization", bearer)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath("/carts")

		// Assertions
		if assert.NoError(t, GetCartController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			// body := rec.Body.String()
			// assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartsWith))
		}
	})

}
