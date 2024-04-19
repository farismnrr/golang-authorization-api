/*
Package middleware provides middleware functions for handling JWT tokens and other authorization-related tasks.

Functions:
- GenerateJWTToken: Generates a JWT token with a short expiration time of 10 seconds.
- ValidateJWTToken: Validates the JWT token.
- IsNonceUsed: Checks if the nonce has been used before.
- AddUsedNonce: Adds the nonce to the used nonces list.
*/

package middleware

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	jwtKey       = []byte(AuthorizationConfig().PrivateKey)
	usedNonces   = make(map[string]bool)
	usedNoncesMu sync.Mutex
)

func GenerateJWTToken() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["exp"] = time.Now().Add(1 * time.Second).Unix()
	claims["nonce"] = fmt.Sprintf("%d", time.Now().UnixNano())

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateJWTToken(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return false, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return false, errors.New("invalid token")
	}

	nonce, ok := claims["nonce"].(string)
	if !ok {
		return false, errors.New("nonce not found in token")
	}

	if IsNonceUsed(nonce) {
		return false, errors.New("nonce has been used before")
	}

	AddUsedNonce(nonce)

	return true, nil
}

func IsNonceUsed(nonce string) bool {
	usedNoncesMu.Lock()
	defer usedNoncesMu.Unlock()
	return usedNonces[nonce]
}

func AddUsedNonce(nonce string) {
	usedNoncesMu.Lock()
	defer usedNoncesMu.Unlock()
	usedNonces[nonce] = true
}
