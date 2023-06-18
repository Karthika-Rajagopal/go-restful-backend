package utils

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// GenerateJWTToken generates a JWT token for the given user ID
func GenerateJWTToken(userID uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte("your-jwt-secret"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateJWTToken validates a JWT token and returns the user ID
func ValidateJWTToken(tokenString string) (uint, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("your-jwt-secret"), nil
	})

	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := uint(claims["userID"].(float64))
		return userID, nil
	}

	return 0, fmt.Errorf("invalid token")
}
