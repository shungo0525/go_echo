package Controllers

import (
	"net/http"

	"github.com/labstack/echo"
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
	{4, "user4", "email4"},
}

func GetUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, users)
}
