package routes

import (
	"github.com/farismnrr/golang-authorization-api/controller"
	"github.com/farismnrr/golang-authorization-api/middleware"
	"github.com/gin-gonic/gin"
)

// SetupRoutes mengatur semua rute aplikasi
func SetupRoutes(router *gin.Engine, controller *controller.CopyrightController) {
	authMiddleware := middleware.AuthorizationMiddleware(middleware.AuthorizationConfig().PrivateKey)

	// Group untuk rute yang memerlukan otorisasi
	router.Use(authMiddleware)

	{
		router.GET("/", controller.GetServer)
		router.GET("/copyright", controller.GetCopyright)
		router.POST("/copyright", controller.AddCopyright)
		router.PUT("/copyright", controller.UpdateCopyright)
		router.DELETE("/copyright", controller.RemoveCopyright)
	}
}
