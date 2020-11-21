package main

import (
	"./Controllers"
	"net/http"
	// "strconv"
	// "fmt"
	// "reflect"  型を確認できる ex)reflect.TypeOf(1)

	"github.com/labstack/echo" // go get -u github.com/labstack/echo/...
)

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var users = []User{
	{1, "user1", "email1"},
	{2, "user2", "email2"},
	{3, "user3", "email3"},
}

func main() {
	e := echo.New()
	e.GET("/", displayHome)
	
	// controllersに切り出し
	e.GET("/users", Controllers.GetUsers)
	e.GET("/users/:id", Controllers.ShowUser)
	e.POST("/users", Controllers.CreateUser)
	e.DELETE("/users/:id", Controllers.DeleteUser)
	e.Logger.Fatal(e.Start(":8080"))

	// e.GET("/users", getUsers)
	// e.GET("/users/:id", showUser)
	// e.POST("/users", createUser)
	// e.GET("/params", getParams)
	// e.DELETE("/users/:id", deleteUser)
}

func displayHome(c echo.Context) error {
	return c.JSON(http.StatusOK, "HOME")
}
