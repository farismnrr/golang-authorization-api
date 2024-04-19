/*
Package middleware provides middleware functions for handling HTTP requests and responses related to Cloudflare API.

Functions:
- GetCloudflare: Sends a GET request to Cloudflare API to retrieve data.
- ParseCloudflareResponse: Parses the response from Cloudflare API.
*/

package middleware

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/farismnrr/golang-authorization-api/model"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func GetCloudflare(ctx *gin.Context) (*model.CloudflareErrorResponse, *model.CloudflareResponse) {
	authToken := AuthorizationConfig()

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	cloudflareApi := os.Getenv("CLOUDFLARE_API_URL")

	req, err := http.NewRequest("GET", cloudflareApi, nil)
	if err != nil {
		errorResponse := &model.CloudflareErrorResponse{
			Success: false,
			Errors: []struct {
				Code    int    `json:"code"`
				Message string `json:"message"`
			}{
				{
					Code:    1001,
					Message: "Failed to create request: " + err.Error(),
				},
			},
			Messages: []interface{}{},
			Result:   nil,
		}
		return errorResponse, nil
	}

	req.Header.Set("Authorization", "Bearer "+authToken.PrivateKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		errorResponse := &model.CloudflareErrorResponse{
			Success: false,
			Errors: []struct {
				Code    int    `json:"code"`
				Message string `json:"message"`
			}{
				{
					Code:    1002,
					Message: "Failed to send request: " + err.Error(),
				},
			},
			Messages: []interface{}{},
			Result:   nil,
		}
		return errorResponse, nil
	}
	defer resp.Body.Close()

	var responseData model.CloudflareResponse
	err = json.NewDecoder(resp.Body).Decode(&responseData)
	if err != nil {
		errorResponse := &model.CloudflareErrorResponse{
			Success: false,
			Errors: []struct {
				Code    int    `json:"code"`
				Message string `json:"message"`
			}{
				{
					Code:    1003,
					Message: "Failed to decode response body: " + err.Error(),
				},
			},
			Messages: []interface{}{},
			Result:   nil,
		}
		return errorResponse, nil
	}

	return nil, &responseData
}

func ParseCloudflareResponse(jsonData []byte) (bool, error) {
	var response model.CloudflareResponse
	err := json.Unmarshal(jsonData, &response)
	if err != nil {
		return false, err
	}
	return response.Success, nil
}
