package Controllers

import (
	"../model"
	"../db"

	"net/http"
	"strconv"
	// "fmt"

	"github.com/labstack/echo"
)

// type User struct {
// 	Id    int    `json:"id"`
// 	Name  string `json:"name"`
// 	Email string `json:"email"`
// }

var users = []model.User{
	{1, "user1", "email1"},
	{2, "user2", "email2"},
	{3, "user3", "email3"},
	{4, "user4", "email4"},
}

// MEMO: function名は大文字でないとだめ。
func GetUsers(c echo.Context) error {
	var users []model.User
	users = db.Index()
	return c.JSON(http.StatusOK, users)
}

func ShowUser(c echo.Context) error {
	var user model.User

	// string -> int
	paramId, _ := strconv.Atoi(c.Param("id"))

	user = db.Show(paramId)

	if user.Id == 0 {
		return echo.NewHTTPError(http.StatusNotFound, "Record Not Found")
	} else {
		return c.JSON(http.StatusOK, user)
	}
}

func CreateUser(c echo.Context) error {
	var user model.User
	user = db.Insert(c.FormValue("name"), c.FormValue("email"))
	
	return c.JSON(http.StatusOK, user)
}

func UpdateUser(c echo.Context) error {
	user := new(model.User)

	// string -> int
	paramId, _ := strconv.Atoi(c.Param("id"))
	for i := 0; i < len(users); i++ {
		if paramId == users[i].Id {
			user = &users[i]
		}
	}
	user.Name = c.FormValue("name")
	user.Email = c.FormValue("email")

	if user.Id == 0 {
		return echo.NewHTTPError(http.StatusNotFound, "Not Found")
	} else {
		return c.JSON(http.StatusOK, user)
	}
}

func DeleteUser(c echo.Context) error {
	var newUsers = []model.User{}
	for i := 0; i < len(users); i++ {
		paramsId, _ := strconv.Atoi(c.Param("id"))
		if users[i].Id != paramsId {
			newUsers = append(newUsers, users[i])
		}
	}

	if (len(newUsers) == len(users)) {
		return echo.NewHTTPError(http.StatusNotFound, "Not Found")
	} else {
		users = newUsers
		return c.JSON(http.StatusOK, newUsers)
	}
}
