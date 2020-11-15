package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", displayHome)
	e.Logger.Fatal(e.Start(":8080"))
}

func displayHome(c echo.Context) error {
	return c.String(http.StatusOK, "HOME")
}
