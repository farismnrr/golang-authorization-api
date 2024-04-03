package controller

import (
	"net/http"

	"github.com/farismnrr/golang-authorization-api/helper"
	"github.com/gin-gonic/gin"
)

// CopyrightController adalah struktur controller untuk mengelola operasi copyright
type CopyrightController struct{}

// GetCopyright digunakan untuk menangani permintaan GET /copyright
func (c *CopyrightController) GetCopyright(ctx *gin.Context) {
	copyrightUsers, err := helper.AuthorizationData()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": http.StatusText(http.StatusInternalServerError),
			"error":   err.Error(),
		})
	}

	if len(copyrightUsers) > 0 {
		// Mengembalikan data hak cipta dalam format JSON
		ctx.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": http.StatusText(http.StatusOK),
			"data":    copyrightUsers,
		})
	} else {
		// Penanganan jika data hak cipta kosong
		ctx.JSON(http.StatusOK, gin.H{
			"status":  http.StatusNoContent,
			"message": http.StatusText(http.StatusNoContent),
			"data":    copyrightUsers,
		})
	}
}
