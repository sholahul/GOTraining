package controller

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateTokenPair(email string) (*TokenPair, error) {
	// Use a secure secret key (replace "your-secret-key" with a strong, random key)
	secretKey := []byte("your-secret-key")

	// Access token
	accessToken := jwt.New(jwt.SigningMethodHS512)
	accessClaims := accessToken.Claims.(jwt.MapClaims)
	accessClaims["email"] = email
	accessClaims["exp"] = time.Now().Add(time.Minute * 1).Unix() // Access token expires in 1 minute

	// Sign the access token
	accessTokenString, err := accessToken.SignedString(secretKey)
	if err != nil {
		return nil, err
	}

	// Refresh token
	refreshToken := jwt.New(jwt.SigningMethodHS512)
	refreshClaims := refreshToken.Claims.(jwt.MapClaims)
	refreshClaims["email"] = email
	refreshClaims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Refresh token expires in 24 hours

	// Sign the refresh token
	refreshTokenString, err := refreshToken.SignedString(secretKey)
	if err != nil {
		return nil, err
	}

	return &TokenPair{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}, nil
}
