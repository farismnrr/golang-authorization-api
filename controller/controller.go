package controller

import (
	"net/http"
	"os"

	"github.com/farismnrr/golang-authorization-api/model"
	"github.com/gin-gonic/gin"
)

// CopyrightController adalah struktur controller untuk mengelola operasi copyright
type CopyrightController struct{}

// CreateCopyright digunakan untuk menangani permintaan POST /copyright
func (c *CopyrightController) CreateCopyright(ctx *gin.Context) {
	// Lakukan sesuatu untuk menangani permintaan POST
	// Di sini, kita hanya akan menampilkan data yang diterima
	var copyright model.Copyright
	if err := ctx.BindJSON(&copyright); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Lakukan sesuatu dengan data copyright (misal: simpan ke database)
	// Di sini, kita hanya akan menampilkan data yang diterima
	ctx.JSON(http.StatusOK, gin.H{"message": "Copyright received", "data": copyright})
}

// GetCopyright digunakan untuk menangani permintaan GET /copyright
func (c *CopyrightController) GetCopyright(ctx *gin.Context) {
	copyright := model.Copyright{
		Username:               os.Getenv("USER_AUTH"),
		CopyrightAuthorization: os.Getenv("AUTHORIZATION"),
	}

	// Mengembalikan data dalam format JSON
	ctx.JSON(http.StatusOK, gin.H{"data": copyright})
}
