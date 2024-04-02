package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthorizationMiddleware adalah middleware untuk otorisasi menggunakan bearer token
func AuthorizationMiddleware(token string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Ambil token dari header Authorization
		bearerToken := c.GetHeader("Authorization")

		// Periksa apakah token diberikan dan benar
		if bearerToken == "" || bearerToken != "Bearer "+token {
			// Jika tidak ada token atau token tidak valid, kirim respons Unauthorized
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Jika token valid, lanjutkan ke handler berikutnya
		c.Next()
	}
}
