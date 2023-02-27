package handler

import (
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/namhoai167/try-echo/models"
)

func Hello(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	username := claims["username"].(string)
	admin := claims["admin"].(bool)

	message := fmt.Sprintf("Hello %s. Admin status: %v", username, admin)

	simpleResponse := &models.SimpleResponse{
		Text: message,
	}

	return c.JSON(http.StatusOK, simpleResponse)
}
