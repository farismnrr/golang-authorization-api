/*
Main function initializes the application:
- Loads environment variables from .env file.
- Initializes a new Gin router.
- Sets up routes for the application.
- Retrieves the port from environment variables or defaults to 8080.
- Runs the server on the specified port.
*/

package main

import (
	"log"
	"os"

	"github.com/farismnrr/golang-authorization-api/controller"
	"github.com/farismnrr/golang-authorization-api/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	router := gin.Default()

	controller := &controller.CopyrightController{}

	routes.SetupRoutes(router, controller, "v1")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":" + port)
}
