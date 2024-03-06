package controller

import (
	"ECHOAPI/model"
	"net/http"

	"github.com/labstack/echo"
)

func GetAllUser(c echo.Context) error {
	var err error
	var data []model.User

	data, err = model.GetAll()

	if err != nil {
		return c.JSON(http.StatusUnauthorized, LoginResponse{
			ResponseCode: "01",
			Message:      "Invalid credentials",
		})
	}

	response := LoginResponse{
		ResponseCode: "00",
		Message:      "Success",
		Data:         data,
	}

	return c.JSON(http.StatusOK, response)
}
