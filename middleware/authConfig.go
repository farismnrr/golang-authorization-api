package middleware

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/farismnrr/golang-authorization-api/model"
	"github.com/gin-gonic/gin"
)

func AuthorizationConfig() model.AuthorizationKey {
	// Inisialisasi middleware untuk otorisasi
	authData, err := ioutil.ReadFile("Authorization.json")
	if err != nil {
		panic(err) // Gagal membaca file, hentikan program
	}

	var auth model.AuthorizationKey
	err = json.Unmarshal(authData, &auth)
	if err != nil {
		panic(err) // Gagal parse JSON, hentikan program
	}

	return auth
}

// AuthorizationMiddleware memeriksa keberadaan dan kevalidan token dalam header Authorization
func AuthorizationMiddleware(token string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Ambil token dari header Authorization
		bearerToken := c.GetHeader("Authorization")

		// Periksa apakah token diberikan
		if bearerToken == "" {
			responseData := model.ResponseStatus{
				Status:  http.StatusForbidden,
				Message: "Access denied: private_key is missing! Failed to Generate JWT Token",
			}
			c.JSON(http.StatusForbidden, responseData)
			c.Abort()
			return
		}

		// Periksa apakah token valid
		if bearerToken != "Bearer "+token {
			// Jika token tidak valid, kirim respons Unauthorized
			responseData := model.ResponseStatus{
				Status:  http.StatusUnauthorized,
				Message: "Access denied: Invalid private_key! Failed to Generate JWT Token",
			}
			c.JSON(http.StatusUnauthorized, responseData)
			c.Abort()
			return
		}

		// Jika token valid, lanjutkan ke handler berikutnya
		c.Next()
	}
}
