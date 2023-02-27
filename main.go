package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/namhoai167/try-echo/handler"
	"github.com/namhoai167/try-echo/mdw"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	isLogedIn := middleware.JWT([]byte(handler.SECRET_KEY))
	e.GET("/", handler.Hello, isLogedIn)
	e.POST("/login", handler.Login, middleware.BasicAuth(mdw.BasicAuth))

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
