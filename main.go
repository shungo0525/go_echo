package main

import (
	"net/http"
	"strconv"
	"./Controllers"
	// "fmt"
	// "reflect"  型を確認できる ex)reflect.TypeOf(1)

	"github.com/labstack/echo" // go get -u github.com/labstack/echo/...
)

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var users = [] User{
	{1, "user1", "email1"},
	{2, "user2", "email2"},
	{3, "user3", "email3"},
}

func main() {
	e := echo.New()
	e.GET("/", displayHome)
	e.GET("/users", Controllers.GetUsers)
	// e.GET("/users", getUsers)
	e.GET("/params", getParams)
	e.GET("/users/:id", showUser)
	e.POST("/users", createUser)
	e.DELETE("/users/:id", deleteUser)
	e.Logger.Fatal(e.Start(":8080"))
}

func displayHome(c echo.Context) error {
	return c.JSON(http.StatusOK, "HOME")
}

func getUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, users)
}

func showUser(c echo.Context) error {
	user := new(User)

	// string -> int
	paramId, _ := strconv.Atoi(c.Param("id"))
	for i := 0; i < len(users); i++ {
		if paramId == users[i].Id {
			user = &users[i]
		}
	}
	return c.JSON(http.StatusOK, user)
}

func getParams(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.JSON(http.StatusOK, name+email)
}

func createUser(c echo.Context) error {
	maxUserId := 0
	for i:=0; i<len(users); i++ {
		if maxUserId < users[i].Id {
			maxUserId = users[i].Id
		}
	}
	newUser := new(User)
	newUser.Id = maxUserId+1
	newUser.Name = c.FormValue("name")
	newUser.Email = c.FormValue("email")

	users = append(users, *newUser)
	return c.JSON(http.StatusOK, users)
}

func deleteUser(c echo.Context) error {
	var newUsers = []User{}
	for i:=0; i<len(users); i++ {
		paramsId, _ := strconv.Atoi(c.Param("id"))
		if users[i].Id != paramsId {
			newUsers = append(newUsers, users[i])
		}
	}
	users = newUsers
	return c.JSON(http.StatusOK, users)
}
