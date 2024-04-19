/*
Package middleware provides middleware functions for handling authorization-related tasks.

Functions:
- AuthorizationConfig: Loads the authorization configuration from the Authorization.json file.
- AuthorizationMiddleware: Middleware function to validate the authorization token.
*/

package middleware

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/farismnrr/golang-authorization-api/model"
	"github.com/gin-gonic/gin"
)

func AuthorizationConfig() model.AuthorizationKey {
	authData, err := ioutil.ReadFile("Authorization.json")
	if err != nil {
		panic(err)
	}

	var auth model.AuthorizationKey
	err = json.Unmarshal(authData, &auth)
	if err != nil {
		panic(err)
	}

	return auth
}

func AuthorizationMiddleware(token string) gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerToken := c.GetHeader("Authorization")

		if bearerToken == "" {
			responseData := model.ResponseStatus{
				Status:  http.StatusForbidden,
				Message: "Access denied: private_key is missing! Failed to Generate JWT Token",
			}
			c.JSON(http.StatusForbidden, responseData)
			c.Abort()
			return
		}

		if bearerToken != "Bearer "+token {
			responseData := model.ResponseStatus{
				Status:  http.StatusUnauthorized,
				Message: "Access denied: Invalid private_key! Failed to Generate JWT Token",
			}
			c.JSON(http.StatusUnauthorized, responseData)
			c.Abort()
			return
		}

		c.Next()
	}
}
