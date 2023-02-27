package mdw

import (
	"log"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func IsAdminMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		admin := claims["admin"].(bool)

		log.Printf("IsAdminMiddleware %v", admin)

		if admin {
			return next(c)
		}

		return echo.ErrUnauthorized
	}
}
