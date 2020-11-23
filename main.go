package main

import (
	// "./Controllers"
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
	e.GET("/users/:id", ControllersGorm.Show)
	e.POST("/users", ControllersGorm.Create)
	e.PUT("/users/:id", ControllersGorm.Update)
	e.DELETE("/users/:id", ControllersGorm.Delete)

	// controllersに切り出し(標準のDB)
	// e.GET("/users", Controllers.GetUsers)
	// e.GET("/users/:id", Controllers.ShowUser)
	// e.POST("/users", Controllers.CreateUser)
	// e.PUT("/users/:id", Controllers.UpdateUser)
	// e.DELETE("/users/:id", Controllers.DeleteUser)

	e.Logger.Fatal(e.Start(":8080"))
}

func displayHome(c echo.Context) error {
	return c.JSON(http.StatusOK, "HOME")
}
