package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type EchoBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	// Create an Echo instance
	e := echo.New()

	// Define routes
	e.POST("/hello", helloHandler)
	e.GET("/greet/:name", greetHandler)
	e.POST("/echo4", echoHandler)

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

func echoHandler(c echo.Context) error {
	// Membaca body dari request
	var echoBody EchoBody
	if err := c.Bind(&echoBody); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Gagal menguraikan JSON"})
	}

	// Melakukan sesuatu dengan data yang diuraikan, misalnya mencetak pesan
	fmt.Printf("Username: %s\n", echoBody.Username)
	fmt.Printf("Password: %s\n", echoBody.Password)
	// Memberikan respon kembali
	// response := map[string]string{"response": "Pesan diterima"}
	response := echoBody
	return c.JSON(http.StatusOK, response)
}
