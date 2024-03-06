package controller

import (
	"ECHOAPI/model"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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

// Use a secure comparison function (e.g., bcrypt.CompareHashAndPassword) in a real-world scenario.
func comparePasswords(hashedPassword, inputPassword string) bool {
	return hashedPassword == inputPassword
}

func LoginHandler(c echo.Context) error {
	// 	// Parse the request body
	var loginReq LoginRequest

	log.Println("Ada yang hit")
	defer c.Bind(&loginReq)

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
	tokenPair, err := GenerateTokenPair(loginReq.Email)
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
	//set cookies
	setTokenCookie(c, accessTokenString)
	setRefreshTokenCookie(c, refreshTokenString)

	return c.JSON(http.StatusOK, response)
}

func Login() {
	// Create an Echo instance
	model.ConnectDB()
	e := echo.New()
	e.Debug = true

	//echo middleware  logger
	// e.Use(middleware.Logger())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "protocol=${protocol}, method=${method}, uri=${uri}, status=${status},error=${error} \n",
	}))

	//BASIC AUTH
	// e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
	// 	// Be careful to use constant time comparison to prevent timing attacks
	// 	user := os.Getenv("AU_USERNAME")
	// 	pass := os.Getenv("AU_PASSWORD")
	// 	if subtle.ConstantTimeCompare([]byte(username), []byte(user)) == 1 &&
	// 		subtle.ConstantTimeCompare([]byte(password), []byte(pass)) == 1 {
	// 		return true, nil
	// 	}
	// 	return false, nil
	// }))

	//SERVER HEADER
	e.Use(ServerHeader)

	//set group
	g := e.Group("/bsb")
	// Define routes
	g.POST("/login", LoginHandler)

	gwt := e.Group("/bsbjwt")
	//use gwt
	gwt.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS512",
		SigningKey:    []byte("bsb"), // Sesuaikan dengan kunci yang benar
		TokenLookup:   "header:Authorization",
		AuthScheme:    "Bearer",
	}))
	gwt.GET("/getalluser", GetAllUser)
	// Start the server
	e.Start(":8080")
}
