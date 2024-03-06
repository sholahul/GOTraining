package router

import (
	"OAUTH2/models"
	"OAUTH2/src/users/usecase"
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	UserRouter struct {
		userUsecase usecase.UsecaseUsers
	}
)

func NewUserRouter(e *echo.Echo) {
	uc := UserRouter{
		userUsecase: usecase.NewUsecaseUser(),
	}
	r := e.Group("/api/users")
	r.GET("", uc.List)
	r.POST("/create", uc.Create)
}

func (h *UserRouter) List(c echo.Context) error {
	var params models.UserParam
	ctx := context.Background()
	err := c.Bind(&params)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	datas, err := h.userUsecase.List(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "ok",
		"code":    http.StatusOK,
		"message": "berhasil mengambil data",
		"data":    datas,
	})
}

func (h *UserRouter) Create(c echo.Context) error {
	var params models.UserParam
	ctx := context.Background()
	err := c.Bind(&params)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = h.userUsecase.Create(ctx, &params)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "ok",
		"code":    http.StatusOK,
		"message": "berhasil membuat data",
	})
}
