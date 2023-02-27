package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	TOKEN    = "RETURNED_TOKEN"
	USERNAME = "admin"
	PASSWORD = "passwd"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)
	e.POST("/login", login, middleware.BasicAuth(basicAuth))

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
	simpleResponse := &SimpleResponse{
		Text: "Hello world",
	}
	return c.JSON(http.StatusOK, simpleResponse)
}

func login(c echo.Context) error {
	return c.JSON(http.StatusOK, &LoginResponse{
		Token: TOKEN,
	})
}

func basicAuth(username, password string, c echo.Context) (bool, error) {
	// Simple verification
	if username != USERNAME || password != PASSWORD {
		return false, nil
	}
	return true, nil
}

// Interface
type LoginResquest struct {
	Username string `json:"username" query:"username" form:"username" xml:"username"`
	Password string `json:"password" query:"password" form:"password" xml:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type SimpleResponse struct {
	Text string `json:"text"`
}
