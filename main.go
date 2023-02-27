package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const TOKEN = "RETURNED_TOKEN"
const USERNAME = "admin"
const PASSWORD = "passwd"

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)
	e.POST("/login", login)

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
	// Init a new struct
	request := new(LoginResquest)

	// Bind the request context to struct
	c.Bind(request)

	// Log
	log.Printf("req data %+v", request)

	// Simple verification
	if request.Username != USERNAME || request.Password != PASSWORD {
		return c.JSON(http.StatusUnauthorized, nil)
	}

	return c.JSON(http.StatusOK, &LoginResponse{
		Token: TOKEN,
	})
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
