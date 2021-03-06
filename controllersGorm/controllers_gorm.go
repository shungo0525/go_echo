package ControllersGorm

import (
	"../gorm_db"
	"../model"

	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func Index(c echo.Context) error {
	var users []model.User
	users = gorm_db.FindAll()
	return c.JSON(http.StatusOK, users)
}

func Show(c echo.Context) error {
	var user model.User
	paramId, _ := strconv.Atoi(c.Param("id"))
	user = gorm_db.Find(paramId)

	if user.Id == 0 {
		return echo.NewHTTPError(http.StatusNotFound, "Record Not Found")
	} else {
		return c.JSON(http.StatusOK, user)
	}
}

func Create(c echo.Context) error {
	var user model.User
	user = gorm_db.Insert(c.FormValue("name"), c.FormValue("email"))

	return c.JSON(http.StatusOK, user)
}

func Update(c echo.Context) error {
	var user model.User
	paramId, _ := strconv.Atoi(c.Param("id"))

	user = gorm_db.Update(paramId, c.FormValue("name"), c.FormValue("email"))

	return c.JSON(http.StatusOK, user)
}

func Delete(c echo.Context) error {
	paramId, _ := strconv.Atoi(c.Param("id"))
	gorm_db.Delete(paramId)

	return c.JSON(http.StatusOK, "Deleted")
}
