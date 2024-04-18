package routes

import (
	"github.com/farismnrr/golang-authorization-api/controller"
	"github.com/farismnrr/golang-authorization-api/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, controller *controller.CopyrightController) {
	{
		router.GET("/get-key", middleware.AuthorizationMiddleware((middleware.AuthorizationConfig().PrivateKey)), controller.ShowCloudflareResponse)

		router.GET("/", controller.GetServer)
		router.GET("/copyright", controller.GetCopyright)
		router.POST("/copyright", controller.AddCopyright)
		router.PUT("/copyright", controller.UpdateCopyright)
		router.DELETE("/copyright", controller.RemoveCopyright)
	}
}
