/*
Package controller provides handlers for various API endpoints related to authorization.

Handlers:
- GetServer: Retrieves server status.
- GetCopyright: Retrieves copyright data.
- AddCopyright: Adds copyright data.
- RemoveCopyright: Removes copyright data.
- UpdateCopyright: Updates copyright data.
- ShowCloudflareResponse: Generates JWT token for Cloudflare authorization.
*/

package controller

import (
	"net/http"

	"github.com/farismnrr/golang-authorization-api/helper"
	"github.com/farismnrr/golang-authorization-api/middleware"
	"github.com/farismnrr/golang-authorization-api/model"
	"github.com/gin-gonic/gin"
)

type CopyrightController struct{}

func (c *CopyrightController) GetServer(ctx *gin.Context) {
	responseStatus := model.ResponseStatus{
		Status:  http.StatusOK,
		Message: "Server is running",
	}
	ctx.JSON(http.StatusOK, responseStatus)
}

func (c *CopyrightController) GetCopyright(ctx *gin.Context) {
	tokenString := ctx.GetHeader("Authorization")
	if tokenString == "" {
		responseData := model.ResponseStatus{
			Status:  http.StatusForbidden,
			Message: "JWT Token is missing",
		}
		ctx.JSON(http.StatusForbidden, responseData)
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
			Message: "JWT Token is missing",
		}
		ctx.JSON(http.StatusForbidden, responseData)
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
		responseData := model.ResponseStatus{
			Status:  http.StatusBadRequest,
			Message: "Failed to add new data",
		}
		ctx.JSON(http.StatusBadRequest, responseData)
		return
	}

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
			Message: "JWT Token is missing",
		}
		ctx.JSON(http.StatusForbidden, responseData)
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

	if !helper.IsUsernameExists(username) {
		responseData := model.ResponseStatus{
			Status:  http.StatusNotFound,
			Message: "Username not found",
		}
		ctx.JSON(http.StatusNotFound, responseData)
		return
	}

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
			Message: "JWT Token is missing",
		}
		ctx.JSON(http.StatusForbidden, responseData)
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

	if !helper.IsUsernameExists(username) {
		responseData := model.ResponseStatus{
			Status:  http.StatusNotFound,
			Message: "Username not found",
		}
		ctx.JSON(http.StatusNotFound, responseData)
		return
	}

	allData, err := helper.UpdateAuthorizationData(requestData)
	if err != nil {
		responseData := model.ResponseStatus{
			Status:  http.StatusBadRequest,
			Message: "Failed to update data: " + err.Error(),
		}
		ctx.JSON(http.StatusBadRequest, responseData)
		return
	}

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
		responseData := model.ResponseStatus{
			Status:  http.StatusBadRequest,
			Message: "Error generating JWT token: " + err.Error(),
		}
		ctx.JSON(http.StatusBadRequest, responseData)
		return
	}

	responseData := model.ResponseStatus{
		Status:  http.StatusAccepted,
		Message: "Access granted: AccessToken JWT generated successfully",
		Data: []model.Copyright{
			{
				CopyrightAuthorization: token,
			},
		},
	}
	ctx.JSON(http.StatusAccepted, responseData)
}
