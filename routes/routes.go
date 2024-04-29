/*
SetupRoutes function sets up the routes for the application.
It takes a Gin router and a CopyrightController instance as parameters.

Routes:
- GET /: Returns server information.
- GET /api/{apiVersion}/get-key: Retrieves the private key for authorization using middleware.
- GET /api/{apiVersion}/copyrights: Retrieves copyright.
- POST /api/{apiVersion}/copyrights: Adds new copyright.
- PUT /api/{apiVersion}/copyrights: Updates existing copyright.
- DELETE /api/{apiVersion}/copyrights: Removes copyright.
*/

package routes

import (
	"fmt"

	"github.com/farismnrr/golang-authorization-api/controller"
	"github.com/farismnrr/golang-authorization-api/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, controller *controller.CopyrightController, apiVersion string) {
	router.GET("/", controller.GetServer)

	apiGroup := router.Group(fmt.Sprintf("/api/%s", apiVersion))
	{
		apiGroup.GET("/get-key", middleware.AuthorizationMiddleware((middleware.AuthorizationConfig().PrivateKey)), controller.ShowCloudflareResponse)
	}

	copyrightGroup := apiGroup.Group("/copyrights")
	{
		copyrightGroup.GET("/", controller.GetCopyright)
		copyrightGroup.POST("/", controller.AddCopyright)
		copyrightGroup.PUT("/", controller.UpdateCopyright)
		copyrightGroup.DELETE("/", controller.RemoveCopyright)
	}
}
