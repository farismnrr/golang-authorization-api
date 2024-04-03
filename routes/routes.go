package routes

import (
	"github.com/farismnrr/golang-authorization-api/controller"
	"github.com/farismnrr/golang-authorization-api/helper"
	"github.com/farismnrr/golang-authorization-api/middleware"
	"github.com/gin-gonic/gin"
)

// SetupRoutes mengatur semua rute aplikasi
func SetupRoutes(router *gin.Engine, controller *controller.CopyrightController) {
	authMiddleware := middleware.AuthorizationMiddleware(helper.AuthorizationKey().PrivateKey)

	// Group untuk rute yang memerlukan otorisasi
	router.Use(authMiddleware)

	{
		router.GET("/copyright", controller.GetCopyright)
	}
}
