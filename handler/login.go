package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/namhoai167/try-echo/models"
)

const (
	TOKEN = "RETURNED_TOKEN"
)

func Login(c echo.Context) error {
	return c.JSON(http.StatusOK, &models.LoginResponse{
		Token: TOKEN,
	})
}
