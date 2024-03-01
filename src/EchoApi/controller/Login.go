package controller

import (
	"ECHOAPI/model"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	ResponseCode string      `json:"responseCode"`
	Message      string      `json:"message"`
	Data         interface{} `json:"data,omitempty"`
}

type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type User struct {
	Email    string
	Password string
	// Add other user properties as needed
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func generateTokenPair(email string) (*TokenPair, error) {
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

// Use a secure comparison function (e.g., bcrypt.CompareHashAndPassword) in a real-world scenario.
func comparePasswords(hashedPassword, inputPassword string) bool {
	return hashedPassword == inputPassword
}

func LoginHandler(c echo.Context) error {
	// 	// Parse the request body
	var loginReq LoginRequest
	if err := c.Bind(&loginReq); err != nil {
		return c.JSON(http.StatusBadRequest, LoginResponse{
			ResponseCode: "99",
			Message:      "Invalid request body",
		})
	}

	// Authenticate the user
	user, err := model.GetUserByEmail(loginReq.Email)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, LoginResponse{
			ResponseCode: "01",
			Message:      "Invalid credentials",
		})
	}

	if !comparePasswords(user.Password, loginReq.Password) {
		return c.JSON(http.StatusUnauthorized, LoginResponse{
			ResponseCode: "01",
			Message:      "Invalid credentials",
		})
	}

	// Authentication successful
	// Generate JWT token
	tokenPair, err := generateTokenPair(loginReq.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, LoginResponse{
			ResponseCode: "99",
			Message:      "Failed to generate token",
		})
	}

	// Convert TokenPair values to strings
	accessTokenString := tokenPair.AccessToken
	refreshTokenString := tokenPair.RefreshToken

	// Authentication successful
	response := LoginResponse{
		ResponseCode: "00",
		Message:      "Success",
		Data: map[string]string{
			"accessToken":  accessTokenString,
			"refreshToken": refreshTokenString,
		},
	}

	return c.JSON(http.StatusOK, response)
}

func Login() {
	// Create an Echo instance
	model.ConnectDB()
	e := echo.New()

	// Define routes
	e.POST("/login", LoginHandler)

	// Start the server
	e.Start(":8080")
}
