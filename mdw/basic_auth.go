package mdw

import "github.com/labstack/echo/v4"

const (
	USERNAME = "admin"
	PASSWORD = "passwd"
)

func BasicAuth(username, password string, c echo.Context) (bool, error) {
	// Simple verification
	if username != USERNAME || password != PASSWORD {
		return false, nil
	}
	return true, nil
}
