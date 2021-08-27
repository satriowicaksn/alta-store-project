package controllers

// func TestGetCartController(t *testing.T) {
// 	t.Run("Test case 1, valid test", func(t *testing.T) {
// 		e := InitEcho()
// 		req := httptest.NewRequest(http.MethodGet, "/", nil)
// 		var token, _ = middlewares.CreateToken(2)
// 		var bearer = "Bearer " + token
// 		req.Header.Add("Authorization", bearer)
// 		rec := httptest.NewRecorder()
// 		c := e.NewContext(req, rec)

// 		c.SetPath("/carts")

// 		// Assertions
// 		if assert.NoError(t, GetCartController(c)) {
// 			assert.Equal(t, http.StatusOK, rec.Code)
// 			// body := rec.Body.String()
// 			// assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartsWith))
// 		}
// 	})

// }
