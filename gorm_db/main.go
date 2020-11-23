package gorm_db

import (
	"../model"
  "github.com/jinzhu/gorm"
  // "strings"
  // "time"
  // "fmt"
  _"github.com/go-sql-driver/mysql"
)

// func main() {
//   db, err := gorm.Open("mysql", "root:@/go_echo")
//   if err != nil {
//     panic(err)
// 	}
// 	defer db.Close()

// 	fmt.Println(findAll(db))
// 	fmt.Println(findById(db, 8))
// 	// fmt.Println(insert(db, "user", "email"))
//   // fmt.Println(update(db, 7, "useruser", "email"))
// }

func FindAll() []model.User {
	db, err := gorm.Open("mysql", "root:@/go_echo")
  if err != nil {
    panic(err)
	}
	defer db.Close()

	var users []model.User
	// ポインタで呼ぶ
	db.Find(&users)
	return users
}

func Find(id int) model.User {
	db, err := gorm.Open("mysql", "root:@/go_echo")
  if err != nil {
    panic(err)
	}
	defer db.Close()

	var user model.User
	// ポインタで呼ぶ
	db.First(&user, id)
	return user
}

func Insert(name string, email string) model.User{
	db, err := gorm.Open("mysql", "root:@/go_echo")
  if err != nil {
    panic(err)
	}
	defer db.Close()

	user := model.User{Name: name, Email: email}
	db.NewRecord(user)
	// ポインタで呼ぶ
	db.Create(&user)
	return user
}

func Update(id int, name string, email string) model.User{
	db, err := gorm.Open("mysql", "root:@/go_echo")
  if err != nil {
    panic(err)
	}
	defer db.Close()

	var user model.User
	db.Model(&user).Where("id = ?", id).Update("name", name, "email", email)
	db.First(&user, id)
	return user
}

func Delete(id int) {
	db, err := gorm.Open("mysql", "root:@/go_echo")
  if err != nil {
    panic(err)
	}
	defer db.Close()

	var user model.User
	db.First(&user, id)
	db.Delete(user)
}

// func deleteByID(id int, db *gorm.DB) {
//     db.Where("id = ?", id).Delete(user)
// }

// deleteByID(1, db)
// fmt.Println(findAll(db)) // []

// https://qiita.com/kai1993/items/389cdec6a01bd527525b
