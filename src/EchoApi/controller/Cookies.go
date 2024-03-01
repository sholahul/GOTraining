package controller

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

var currentTime time.Time
var err error

func init_date() {
	currentTime, err = Get_Date()
	if err != nil {
		// Handle the error (e.g., log it)
		fmt.Printf("Error getting date: %s\n", err)
	}
}

func setTokenCookie(c echo.Context, token string) error {
	Strtimes := os.Getenv("TIME_ACCESS_TOKEN")
	times, _ := strconv.Atoi(Strtimes)
	init_date()

	cookie := new(http.Cookie)
	cookie.Name = "access_token"
	cookie.Value = token
	cookie.Expires = currentTime.Add(time.Duration(times) * time.Minute) // Sesuaikan dengan kebutuhan Anda
	log.Println("Access_Token Expire : ", cookie.Expires)
	c.SetCookie(cookie)
	return nil
}

func setRefreshTokenCookie(c echo.Context, refreshToken string) error {
	Strtimes := os.Getenv("TIME_REFRESH_TOKEN")
	times, _ := strconv.Atoi(Strtimes)
	init_date()

	cookie := new(http.Cookie)
	cookie.Name = "refresh_token"
	cookie.Value = refreshToken
	cookie.Expires = currentTime.Add(time.Duration(times) * time.Minute) // Sesuaikan dengan kebutuhan Anda
	log.Println("Refresh_Token Expire : ", cookie.Expires)

	c.SetCookie(cookie)
	return nil
}
