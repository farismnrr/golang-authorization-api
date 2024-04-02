package routes

import (
	"github.com/farismnrr/golang-authorization-api/controller"
	"github.com/farismnrr/golang-authorization-api/middleware"
	"github.com/gin-gonic/gin"
)

// SetupRoutes mengatur semua rute aplikasi
func SetupRoutes(router *gin.Engine, controller *controller.CopyrightController) {
	// Inisialisasi middleware untuk otorisasi
	authMiddleware := middleware.AuthorizationMiddleware("copyright-handled-by-farismnrr")

	// Group untuk rute yang memerlukan otorisasi
	router.Use(authMiddleware)

	{
		router.GET("/copyright", controller.GetCopyright)
	}
}
