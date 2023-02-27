package handler

import (
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/namhoai167/try-echo/models"
)

const (
	SECRET_KEY = "secret_key"
)

func Login(c echo.Context) error {

	username := c.Get("username").(string)
	admin := c.Get("admin").(bool)
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["admin"] = admin
	claims["exp"] = time.Now().Add(2 * time.Minute).Unix()

	t, err := token.SignedString([]byte(SECRET_KEY))

	if err != nil {
		log.Printf("Signed token err %v\n", err)
		return err
	}

	return c.JSON(http.StatusOK, &models.LoginResponse{
		Token: t,
	})
}
