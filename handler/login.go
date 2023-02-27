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
	// Get the username, admin status from request
	username := c.Get("username").(string)
	admin := c.Get("admin").(bool)

	// Construct JWT with username, admin status and expired time
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["admin"] = admin
	claims["exp"] = time.Now().Add(2 * time.Minute).Unix()

	// Sign JWT with secret key
	t, err := token.SignedString([]byte(SECRET_KEY))

	// Handle error
	if err != nil {
		log.Printf("Signed token err %v\n", err)
		return err
	}

	// Return JWT
	return c.JSON(http.StatusOK, &models.LoginResponse{
		Token: t,
	})
}
