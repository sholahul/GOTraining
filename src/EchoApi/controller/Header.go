package controller

import "github.com/labstack/echo"

func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "BSBLINK V.0")
		//custom middleware
		c.Response().Header().Set("Developer", "TIM IT SOLUTION 2024")
		c.Response().Header().Set("Date", Get_date())
		return next(c)
	}
}
