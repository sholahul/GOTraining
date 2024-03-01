package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	// fmt.Println("hello world")
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello world")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
