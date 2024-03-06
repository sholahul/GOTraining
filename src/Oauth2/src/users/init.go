package src

import (
	"OAUTH2/src/users/router"

	"github.com/labstack/echo/v4"
)

func InitRouter(e *echo.Echo) {
	router.NewUserRouter(e)
}
