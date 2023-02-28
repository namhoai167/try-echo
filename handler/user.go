package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUser(c echo.Context) error {
	return c.String(http.StatusOK, "GET user")
}

func UpdateUser(c echo.Context) error {
	return c.String(http.StatusOK, "PUT user")
}

func DeleteUser(c echo.Context) error {
	return c.String(http.StatusOK, "DELETE user")
}
