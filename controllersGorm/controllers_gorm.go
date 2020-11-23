package ControllersGorm

import (
	"../model"
	"../gorm_db"

	"net/http"
	// "strconv"
	"fmt"

	"github.com/labstack/echo"
)

func Index(c echo.Context) error {
	var users []model.User
	users = gorm_db.FindAll()
	fmt.Println(users)
	return c.JSON(http.StatusOK, users)
}
