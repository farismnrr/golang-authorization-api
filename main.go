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
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Inisialisasi router
	router := gin.Default()

	// Inisialisasi controller
	controller := &controller.CopyrightController{}

	// Setup routes
	routes.SetupRoutes(router, controller)

	// Mulai server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port
	}
	router.Run(":" + port)
}
