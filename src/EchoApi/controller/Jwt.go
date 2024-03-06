package controller

import (
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateTokenPair(email string) (*TokenPair, error) {
	time_access_token := os.Getenv("TIME_ACCESS_TOKEN")
	times_acc, _ := strconv.Atoi(time_access_token)
	time_refresh_token := os.Getenv("TIME_REFRESH_TOKEN")
	times_ref, _ := strconv.Atoi(time_refresh_token)

	init_date()

	// Use a secure secret key (replace "your-secret-key" with a strong, random key)
	secretKey := []byte("bsb")

	// Access token
	accessToken := jwt.New(jwt.SigningMethodHS512)
	accessClaims := accessToken.Claims.(jwt.MapClaims)
	accessClaims["email"] = email
	accessClaims["exp"] = currentTime.Add(time.Duration(times_acc) * time.Minute) // Access token expires in 1 minute

	// Sign the access token
	accessTokenString, err := accessToken.SignedString(secretKey)
	if err != nil {
		return nil, err
	}

	// Refresh token

	refreshToken := jwt.New(jwt.SigningMethodHS512)
	refreshClaims := refreshToken.Claims.(jwt.MapClaims)
	refreshClaims["email"] = email
	refreshClaims["exp"] = currentTime.Add(time.Duration(times_ref) * time.Minute) // Refresh token expires in 24 hours

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
