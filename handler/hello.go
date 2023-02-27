package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/namhoai167/try-echo/models"
)

func Hello(c echo.Context) error {
	simpleResponse := &models.SimpleResponse{
		Text: "Hello world!",
	}
	return c.JSON(http.StatusOK, simpleResponse)
}
