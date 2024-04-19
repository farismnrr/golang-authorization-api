/*
Package model provides structures for storing various data related to the application.

Structures:
- Copyright: Stores data related to copyright.
- AuthorizationKey: Stores the private key for authorization.
- ResponseStatus: Stores status and message of the response.
- CloudflareResponse: Represents the response from Cloudflare API.
- CloudflareErrorResponse: Represents the error response from Cloudflare API.
*/

package model

type Copyright struct {
	Id                     string `json:"id,omitempty"`
	Username               string `json:"username,omitempty"`
	NewUsername            string `json:"newUsername,omitempty"`
	CopyrightAuthorization string `json:"copyrightAuthorization"`
}

type AuthorizationKey struct {
	PrivateKey string `json:"private_key"`
}

type ResponseStatus struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    []Copyright `json:"data,omitempty"`
}

type CloudflareResponse struct {
	Errors   []interface{} `json:"errors"`
	Messages []struct {
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Type    interface{} `json:"type"`
	} `json:"messages"`
	Result struct {
		ID     string `json:"id"`
		Status string `json:"status"`
	} `json:"result"`
	Success bool `json:"success"`
}

type CloudflareErrorResponse struct {
	Success bool `json:"success"`
	Errors  []struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"errors"`
	Messages []interface{} `json:"messages"`
	Result   interface{}   `json:"result"`
}
