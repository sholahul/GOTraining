package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	// Create an Echo instance
	e := echo.New()

	// Define routes
	e.POST("/hello", helloHandler)
	e.GET("/greet/:name", greetHandler)

	// Start the server
	e.Start(":8080")
}

// Handler for /hello route
func helloHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

// Handler for /greet/:name route
func greetHandler(c echo.Context) error {
	name := c.Param("name")
	return c.String(http.StatusOK, "Hello, "+name+"!")
}
