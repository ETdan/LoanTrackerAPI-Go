package infrastructure

import (
	domain "LoanAPI/LoanTrackerAPI-Go/Domain"
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtSecret = []byte("your_secret_key") // Replace with your secret key

// GenerateJWTToken generates a JWT token for email verification
func GenerateJWTToken(user domain.UnverifiedUser) (string, error) {
	claims := jwt.MapClaims{
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour).Unix(), // Token expires in 24 hours
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
func GenerateUserJWTToken(user domain.User) (string, error) {
	claims := jwt.MapClaims{
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour).Unix(), // Token expires in 24 hours
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func GenerateAccessTokenAndRefreshToken(user_id string) (string, string, error) {
	accessclaims := jwt.MapClaims{
		"user_id": user_id,
		"exp":     time.Now().Add(time.Hour).Unix(), // Token expires in 24 hours
	}

	refreshclaims := jwt.MapClaims{
		"user_id": user_id,
		"exp":     time.Now().Add(7 * 24 * time.Hour).Unix(), // Token expires in 24 hours
	}

	acesstoken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessclaims)
	refreshtoken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshclaims)

	accessToken, err := acesstoken.SignedString(jwtSecret)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := refreshtoken.SignedString(jwtSecret)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
