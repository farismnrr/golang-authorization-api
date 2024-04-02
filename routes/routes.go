package routes

import (
	"encoding/json"
	"io/ioutil"

	"github.com/farismnrr/golang-authorization-api/controller"
	"github.com/farismnrr/golang-authorization-api/middleware"
	"github.com/farismnrr/golang-authorization-api/model"
	"github.com/gin-gonic/gin"
)

// SetupRoutes mengatur semua rute aplikasi
func SetupRoutes(router *gin.Engine, controller *controller.CopyrightController) {
	// Inisialisasi middleware untuk otorisasi
	authData, err := ioutil.ReadFile("Authorization.json")
	if err != nil {
		panic(err) // Gagal membaca file, hentikan program
	}

	var auth model.AuthorizationData
	err = json.Unmarshal(authData, &auth)
	if err != nil {
		panic(err) // Gagal parse JSON, hentikan program
	}

	authMiddleware := middleware.AuthorizationMiddleware(auth.PrivateKey)

	// Group untuk rute yang memerlukan otorisasi
	router.Use(authMiddleware)

	{
		router.GET("/copyright", controller.GetCopyright)
	}
}
