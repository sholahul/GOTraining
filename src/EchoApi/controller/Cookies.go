package controller

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

func setTokenCookie(c echo.Context, token string) error {
	Strtimes := os.Getenv("TIME_ACCESS_TOKEN")
	times, _ := strconv.Atoi(Strtimes)

	cookie := new(http.Cookie)
	cookie.Name = "access_token"
	cookie.Value = token
	cookie.Expires = time.Now().Add(time.Duration(times) * time.Minute) // Sesuaikan dengan kebutuhan Anda
	log.Println("Duration Access token : ", time.Duration(times))
	c.SetCookie(cookie)
	return nil
}

func setRefreshTokenCookie(c echo.Context, refreshToken string) error {
	Strtimes := os.Getenv("TIME_REFRESH_TOKEN")
	times, _ := strconv.Atoi(Strtimes)

	cookie := new(http.Cookie)
	cookie.Name = "refresh_token"
	cookie.Value = refreshToken
	cookie.Expires = time.Now().Add(time.Duration(times) * time.Hour) // Sesuaikan dengan kebutuhan Anda
	log.Println("Duration Refresh token : ", time.Duration(times))
	c.SetCookie(cookie)
	return nil
}
