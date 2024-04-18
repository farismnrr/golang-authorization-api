package controller

import (
	"net/http"

	"github.com/farismnrr/golang-authorization-api/helper"
	"github.com/farismnrr/golang-authorization-api/middleware"
	"github.com/farismnrr/golang-authorization-api/model"
	"github.com/gin-gonic/gin"
)

// CopyrightController adalah struktur controller untuk mengelola operasi copyright
type CopyrightController struct{}

// Fungsi ini akan menangani permintaan GET ke endpoint "/"
func (c *CopyrightController) GetServer(ctx *gin.Context) {
	tokenString := ctx.GetHeader("Authorization")
	if tokenString == "" {
		responseData := model.ResponseStatus{
			Status:  http.StatusForbidden,
			Message: "Authorization header is missing",
		}
		ctx.JSON(http.StatusBadRequest, responseData)
		return
	}

	tokenString = tokenString[7:]

	_, err := middleware.ValidateJWTToken(tokenString)
	if err != nil {
		responseData := model.ResponseStatus{
			Status:  http.StatusUnauthorized,
			Message: "Invalid token",
		}
		ctx.JSON(http.StatusUnauthorized, responseData)
		return
	}

	responseStatus := model.ResponseStatus{
		Status:  http.StatusOK,
		Message: "Server is running",
	}
	ctx.JSON(http.StatusOK, responseStatus)
}

// GetCopyright digunakan untuk menangani permintaan GET /copyright
func (c *CopyrightController) GetCopyright(ctx *gin.Context) {
	tokenString := ctx.GetHeader("Authorization")
	if tokenString == "" {
		responseData := model.ResponseStatus{
			Status:  http.StatusForbidden,
			Message: "Authorization header is missing",
		}
		ctx.JSON(http.StatusBadRequest, responseData)
		return
	}

	tokenString = tokenString[7:]

	_, err := middleware.ValidateJWTToken(tokenString)
	if err != nil {
		responseData := model.ResponseStatus{
			Status:  http.StatusUnauthorized,
			Message: "Invalid token",
		}
		ctx.JSON(http.StatusUnauthorized, responseData)
		return
	}

	helper.AddDummyAuthorizationData()
	copyrightUsers, err := helper.AuthorizationData()

	if err != nil {
		responseStatus := model.ResponseStatus{
			Status:  http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
		}

		ctx.JSON(http.StatusInternalServerError, responseStatus)
		return
	}

	if len(copyrightUsers) > 0 {
		responseData := model.ResponseStatus{
			Status:  http.StatusOK,
			Message: http.StatusText(http.StatusOK),
			Data:    copyrightUsers,
		}

		ctx.JSON(http.StatusOK, responseData)
		return
	} else {
		responseData := model.ResponseStatus{
			Status:  http.StatusNoContent,
			Message: http.StatusText(http.StatusNoContent),
			Data:    copyrightUsers,
		}

		ctx.JSON(http.StatusNoContent, responseData)
		return
	}
}

func (c *CopyrightController) AddCopyright(ctx *gin.Context) {
	tokenString := ctx.GetHeader("Authorization")
	if tokenString == "" {
		responseData := model.ResponseStatus{
			Status:  http.StatusForbidden,
			Message: "Authorization header is missing",
		}
		ctx.JSON(http.StatusBadRequest, responseData)
		return
	}

	tokenString = tokenString[7:]

	_, err := middleware.ValidateJWTToken(tokenString)
	if err != nil {
		responseData := model.ResponseStatus{
			Status:  http.StatusUnauthorized,
			Message: "Invalid token",
		}
		ctx.JSON(http.StatusUnauthorized, responseData)
		return
	}

	// Membaca body request
	var requestData map[string]string

	if err := ctx.BindJSON(&requestData); err != nil {
		responseData := model.ResponseStatus{
			Status:  http.StatusBadRequest,
			Message: "Failed to read request body",
		}

		ctx.JSON(http.StatusBadRequest, responseData)
		return
	}

	username, ok := requestData["username"]
	if !ok || username == "" {
		responseData := model.ResponseStatus{
			Status:  http.StatusBadRequest,
			Message: "Username is required",
		}

		ctx.JSON(http.StatusBadRequest, responseData)
		return
	}

	// Memastikan hanya key "username" yang diperbolehkan
	for key := range requestData {
		if key != "username" {
			responseData := model.ResponseStatus{
				Status:  http.StatusBadRequest,
				Message: "Key " + key + " is not allowed",
			}
			ctx.JSON(http.StatusBadRequest, responseData)
			return // Pindahkan return ke sini agar bisa menangani semua key
		}
	}

	// Mengecek apakah username sudah ada
	if helper.IsUsernameExists(username) {
		responseData := model.ResponseStatus{
			Status:  http.StatusBadRequest,
			Message: "Username already exists",
		}
		ctx.JSON(http.StatusBadRequest, responseData)
		return
	}

	newData, err := helper.AddAuthorizationData(username)
	if err != nil {
		// Handle error jika gagal menambahkan data
		responseData := model.ResponseStatus{
			Status:  http.StatusBadRequest,
			Message: "Failed to add new data",
		}
		ctx.JSON(http.StatusBadRequest, responseData)
		return
	}

	// Handle data (misalnya simpan ke database, dll.)
	// Kemudian bisa mengembalikan response, misalnya:
	responseData := model.ResponseStatus{
		Status:  http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data:    newData,
	}
	ctx.JSON(http.StatusOK, responseData)
}

func (c *CopyrightController) RemoveCopyright(ctx *gin.Context) {
	tokenString := ctx.GetHeader("Authorization")
	if tokenString == "" {
		responseData := model.ResponseStatus{
			Status:  http.StatusForbidden,
			Message: "Authorization header is missing",
		}
		ctx.JSON(http.StatusBadRequest, responseData)
		return
	}

	tokenString = tokenString[7:]

	_, err := middleware.ValidateJWTToken(tokenString)
	if err != nil {
		responseData := model.ResponseStatus{
			Status:  http.StatusUnauthorized,
			Message: "Invalid token",
		}
		ctx.JSON(http.StatusUnauthorized, responseData)
		return
	}

	// Membaca body request
	var requestData map[string]string

	if err := ctx.BindJSON(&requestData); err != nil {
		responseData := model.ResponseStatus{
			Status:  http.StatusBadRequest,
			Message: "Failed to read request body",
		}
		ctx.JSON(http.StatusBadRequest, responseData)
		return
	}

	username, ok := requestData["username"]
	if !ok || username == "" {
		responseData := model.ResponseStatus{
			Status:  http.StatusBadRequest,
			Message: "Username is required",
		}
		ctx.JSON(http.StatusBadRequest, responseData)
		return
	}

	// Memastikan hanya key "username" yang diperbolehkan
	for key := range requestData {
		if key != "username" {
			responseData := model.ResponseStatus{
				Status:  http.StatusBadRequest,
				Message: "Key " + key + " is not allowed",
			}
			ctx.JSON(http.StatusBadRequest, responseData)
			return
		}
	}

	// Mengecek apakah username ada dalam data
	if !helper.IsUsernameExists(username) {
		responseData := model.ResponseStatus{
			Status:  http.StatusNotFound,
			Message: "Username not found",
		}
		ctx.JSON(http.StatusNotFound, responseData)
		return
	}

	// Menghapus data berdasarkan username
	deletedData, err := helper.RemoveAuthorizationData(username)
	if err != nil {
		responseData := model.ResponseStatus{
			Status:  http.StatusInternalServerError,
			Message: "Failed to delete data",
		}
		ctx.JSON(http.StatusInternalServerError, responseData)
		return
	}

	responseData := model.ResponseStatus{
		Status:  http.StatusOK,
		Message: "Data deleted successfully",
		Data:    deletedData,
	}
	ctx.JSON(http.StatusOK, responseData)
}

func (c *CopyrightController) UpdateCopyright(ctx *gin.Context) {
	tokenString := ctx.GetHeader("Authorization")
	if tokenString == "" {
		responseData := model.ResponseStatus{
			Status:  http.StatusForbidden,
			Message: "Authorization header is missing",
		}
		ctx.JSON(http.StatusBadRequest, responseData)
		return
	}

	tokenString = tokenString[7:]

	_, err := middleware.ValidateJWTToken(tokenString)
	if err != nil {
		responseData := model.ResponseStatus{
			Status:  http.StatusUnauthorized,
			Message: "Invalid token",
		}
		ctx.JSON(http.StatusUnauthorized, responseData)
		return
	}

	// Membaca body request
	var requestData map[string]string

	if err := ctx.BindJSON(&requestData); err != nil {
		responseData := model.ResponseStatus{
			Status:  http.StatusBadRequest,
			Message: "Failed to read request body",
		}
		ctx.JSON(http.StatusBadRequest, responseData)
		return
	}

	username, ok := requestData["username"]
	if !ok || username == "" {
		responseData := model.ResponseStatus{
			Status:  http.StatusBadRequest,
			Message: "Username is required",
		}
		ctx.JSON(http.StatusBadRequest, responseData)
		return
	}

	// Memastikan hanya key "username" yang diperbolehkan
	for key := range requestData {
		if (key != "username") && (key != "newUsername") {
			responseData := model.ResponseStatus{
				Status:  http.StatusBadRequest,
				Message: "Key " + key + " is not allowed",
			}
			ctx.JSON(http.StatusBadRequest, responseData)
			return
		}
	}

	// Mengecek apakah username ada dalam data
	if !helper.IsUsernameExists(username) {
		responseData := model.ResponseStatus{
			Status:  http.StatusNotFound,
			Message: "Username not found",
		}
		ctx.JSON(http.StatusNotFound, responseData)
		return
	}

	// Mengupdate data
	allData, err := helper.UpdateAuthorizationData(requestData)
	if err != nil {
		// Handle error jika gagal mengupdate data
		responseData := model.ResponseStatus{
			Status:  http.StatusBadRequest,
			Message: "Failed to update data: " + err.Error(),
		}
		ctx.JSON(http.StatusBadRequest, responseData)
		return
	}

	// Handle data (misalnya simpan ke database, dll.)
	// Kemudian bisa mengembalikan response, misalnya:
	var responseData model.ResponseStatus
	if len(allData) > 0 {
		responseData = model.ResponseStatus{
			Status:  http.StatusOK,
			Message: http.StatusText(http.StatusOK),
			Data: []model.Copyright{
				{
					Id:                     allData[0].Id,
					Username:               requestData["username"],
					NewUsername:            allData[0].Username,
					CopyrightAuthorization: allData[0].CopyrightAuthorization,
				},
			},
		}
	}
	ctx.JSON(http.StatusOK, responseData)
}

func (c *CopyrightController) ShowCloudflareResponse(ctx *gin.Context) {
	token, err := middleware.GenerateJWTToken()

	if err != nil {
		// Handle error jika gagal mengupdate data
		responseData := model.ResponseStatus{
			Status:  http.StatusBadRequest,
			Message: "Error generating JWT token: " + err.Error(),
		}
		ctx.JSON(http.StatusBadRequest, responseData)
		return
	}

	// Handle error jika gagal mengupdate data
	responseData := model.ResponseStatus{
		Status:  http.StatusAccepted,
		Message: "Token JWT Generated Succesfully",
		Data: []model.Copyright{
			{
				CopyrightAuthorization: token,
			},
		},
	}
	ctx.JSON(http.StatusAccepted, responseData)
}
