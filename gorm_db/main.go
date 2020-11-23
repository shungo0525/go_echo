package main

import (
	"../model"
  "github.com/jinzhu/gorm"
  // "strings"
  // "time"
  "fmt"
  _"github.com/go-sql-driver/mysql"
)

func main() {
  db, err := gorm.Open("mysql", "root:@/go_echo")

  if err != nil {
    panic(err)
	}

	defer db.Close()

	fmt.Println(findAll(db))
	fmt.Println(findById(db, 8))

}

func findAll(db *gorm.DB) []model.User {
	var allUsers []model.User
	db.Find(&allUsers)
	return allUsers
}

func findById(db *gorm.DB, id int) model.User {
	var user model.User
	db.First(&user, id)
	return user
}

// https://qiita.com/kai1993/items/389cdec6a01bd527525b
