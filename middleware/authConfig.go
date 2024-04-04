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

// AuthorizationMiddleware adalah middleware untuk otorisasi menggunakan bearer token
func AuthorizationMiddleware(token string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Ambil token dari header Authorization
		bearerToken := c.GetHeader("Authorization")

		// Periksa apakah token diberikan dan benar
		if bearerToken == "" || bearerToken != "Bearer "+token {
			// Jika tidak ada token atau token tidak valid, kirim respons Unauthorized
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"message": http.StatusText(http.StatusUnauthorized),
			})
			c.Abort()
			return
		}

		// Jika token valid, lanjutkan ke handler berikutnya
		c.Next()
	}
}
