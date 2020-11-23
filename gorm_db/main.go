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
	// fmt.Println(insert(db, "user", "email"))
  fmt.Println(update(db, 7, "useruser", "email"))
}

func findAll(db *gorm.DB) []model.User {
	var allUsers []model.User
	// ポインタで呼ぶ
	db.Find(&allUsers)
	return allUsers
}

func findById(db *gorm.DB, id int) model.User {
	var user model.User
	// ポインタで呼ぶ
	db.First(&user, id)
	return user
}

func insert(db *gorm.DB, name string, email string) model.User{
	user := model.User{Name: name, Email: email}
	db.NewRecord(user)
	// ポインタで呼ぶ
	db.Create(&user)
	return user
}

func update(db *gorm.DB, id int, name string, email string) model.User{
	var user model.User
	db.Model(&user).Where("id = ?", id).Update("name", name, "email", email)
	db.First(&user, id)
	return user
}

// https://qiita.com/kai1993/items/389cdec6a01bd527525b
