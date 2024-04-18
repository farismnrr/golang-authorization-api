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
	usedNonces   = make(map[string]bool) // Map to store used nonces
	usedNoncesMu sync.Mutex              // Mutex for concurrent access to usedNonces
)

// GenerateJWTToken generates a JWT token with a short expiration time of 10 seconds
func GenerateJWTToken() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["exp"] = time.Now().Add(10 * time.Minute).Unix()
	claims["nonce"] = fmt.Sprintf("%d", time.Now().UnixNano()) // Use current timestamp as nonce

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateJWTToken validates the JWT token
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

	// Check if nonce exists
	nonce, ok := claims["nonce"].(string)
	if !ok {
		return false, errors.New("nonce not found in token")
	}

	// Check if nonce has been used before
	if IsNonceUsed(nonce) {
		return false, errors.New("nonce has been used before")
	}

	// Mark nonce as used
	AddUsedNonce(nonce)

	return true, nil
}

// IsNonceUsed checks if the nonce has been used before
func IsNonceUsed(nonce string) bool {
	usedNoncesMu.Lock()
	defer usedNoncesMu.Unlock()
	return usedNonces[nonce]
}

// AddUsedNonce adds the nonce to the used nonces list
func AddUsedNonce(nonce string) {
	usedNoncesMu.Lock()
	defer usedNoncesMu.Unlock()
	usedNonces[nonce] = true
}
