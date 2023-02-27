package mdw

import "github.com/labstack/echo/v4"

const (
	USERNAME = "admin"
	PASSWORD = "passwd"
)

func BasicAuth(username, password string, c echo.Context) (bool, error) {
	// Simple verification
	if username == USERNAME && password == PASSWORD {
		c.Set("username", username)
		c.Set("admin", true)
		return true, nil
	}

	if username == "nam" && password == PASSWORD {
		c.Set("username", "nam")
		c.Set("admin", false)
		return true, nil
	}

	return false, nil
}
