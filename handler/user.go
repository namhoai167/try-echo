package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

var listUser = []User{
	{Name: "Noah", Age: 17},
	{Name: "Jack", Age: 18},
	{Name: "Carter", Age: 19},
	{Name: "Evan", Age: 20},
	{Name: "Evan", Age: 21},
	{Name: "Frank", Age: 22},
	{Name: "Hazel", Age: 23},
	{Name: "Chloe", Age: 24},
	{Name: "Janny", Age: 25},
	{Name: "Jason", Age: 26},
	{Name: "Judia", Age: 27},
	{Name: "Liam", Age: 28},
	{Name: "Lizzie", Age: 29},
	{Name: "Mike", Age: 30},
	{Name: "Philip", Age: 31},
	{Name: "Wolfe", Age: 32},
	{Name: "Ben", Age: 33},
	{Name: "Halie", Age: 34},
	{Name: "Harry", Age: 35},
	{Name: "Luis", Age: 36},
	{Name: "Paul", Age: 37},
	{Name: "Roy", Age: 38},
	{Name: "Stella", Age: 39},
}

func GetUser(c echo.Context) error {
	return c.String(http.StatusOK, "GET user")
}

func UpdateUser(c echo.Context) error {
	return c.String(http.StatusOK, "PUT user")
}

func DeleteUser(c echo.Context) error {
	return c.String(http.StatusOK, "DELETE user")
}

func GetAllUsers(c echo.Context) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	c.Response().WriteHeader(http.StatusOK)

	enc := json.NewEncoder(c.Response())

	for _, user := range listUser {
		if err := enc.Encode(user); err != nil {
			return err
		}
		c.Response().Flush()
		time.Sleep(1 * time.Second)
	}

	return nil
}

type User struct {
	Name string
	Age  int
}
