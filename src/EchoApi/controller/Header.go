package controller

import (
	"fmt"
	"time"

	"github.com/labstack/echo"
)

func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	currentTime, err := Get_Date()
	if err != nil {
		// Handle the error (e.g., log it)
		fmt.Printf("Error getting date: %s\n", err)
	}

	// Format the time as per your requirement (e.g., RFC1123)
	formattedTime := currentTime.Format(time.RFC1123)

	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "BSBLINK V.0")
		//custom middleware
		c.Response().Header().Set("Developer", "TIM IT SOLUTION 2024")
		c.Response().Header().Set("Date", formattedTime)
		return next(c)
	}
}
