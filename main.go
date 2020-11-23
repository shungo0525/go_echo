package main

import (
	"./Controllers"
	"./ControllersGorm"
	"net/http"
	// "strconv"
	// "fmt"
	// "reflect"  型を確認できる ex)reflect.TypeOf(1)

	"github.com/labstack/echo" // go get -u github.com/labstack/echo/...
)

func main() {
	e := echo.New()
	e.GET("/", displayHome)

	// controllersGormに切り出し
	e.GET("/users", ControllersGorm.Index)

	// controllersに切り出し
	// e.GET("/users", Controllers.GetUsers)
	e.GET("/users/:id", Controllers.ShowUser)
	e.POST("/users", Controllers.CreateUser)
	e.PUT("/users/:id", Controllers.UpdateUser)
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
