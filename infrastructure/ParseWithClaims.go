package infrastructure

import (
	"fmt"

	"github.com/golang-jwt/jwt"
)

func VerifyRefreshToken(refresh_token string) (string, error) {
	// Parse the token
	token, err := jwt.Parse(refresh_token, func(token *jwt.Token) (interface{}, error) {
		// Ensure that the signing method is what you expect
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	if err != nil {
		return "", err
	}

	// Check if the token is valid and extract claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, ok := claims["user_id"].(string)
		if !ok {
			return "", fmt.Errorf("user_id not found in token")
		}
		return userID, nil
	}

	return "", fmt.Errorf("invalid token")
}

func ParseClaims(tokenString string) (string, error) {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil || !token.Valid {
		return "", err
	}

	email, ok := claims["email"].(string)
	if !ok {
		return "", err
	}

	return email, nil
}
