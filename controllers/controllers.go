package Controllers

import (
	"../model"
	"../db"

	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

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
	var user model.User

	// string -> int
	paramId, _ := strconv.Atoi(c.Param("id"))

	user = db.Update(paramId, c.FormValue("name"), c.FormValue("email"))

	if user.Id == 0 {
		return echo.NewHTTPError(http.StatusNotFound, "Not Found")
	} else {
		return c.JSON(http.StatusOK, user)
	}
}

func DeleteUser(c echo.Context) error {
	var user model.User
	paramId, _ := strconv.Atoi(c.Param("id"))

	user = db.Delete(paramId)

	if user.Id == 0 {
		return echo.NewHTTPError(http.StatusNotFound, "Not Found")
	} else {
		return c.JSON(http.StatusOK, user)
	}
}
