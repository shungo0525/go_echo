package Controllers

import (
	"../model"
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
	return c.JSON(http.StatusOK, users)
}

func ShowUser(c echo.Context) error {
	user := new(model.User)

	// string -> int
	paramId, _ := strconv.Atoi(c.Param("id"))
	for i := 0; i < len(users); i++ {
		if paramId == users[i].Id {
			user = &users[i]
		}
	}

	if user.Id == 0 {
		return echo.NewHTTPError(http.StatusNotFound, "Not Found")
	} else {
		return c.JSON(http.StatusOK, user)
	}
}

func CreateUser(c echo.Context) error {
	maxUserId := 0
	for i := 0; i < len(users); i++ {
		if maxUserId < users[i].Id {
			maxUserId = users[i].Id
		}
	}
	newUser := new(model.User)
	newUser.Id = maxUserId + 1
	newUser.Name = c.FormValue("name")
	newUser.Email = c.FormValue("email")

	users = append(users, *newUser)
	return c.JSON(http.StatusOK, users)
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
