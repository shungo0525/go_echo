package gorm_db

import (
	"../model"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func initDB() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@/go_echo")

	if err != nil {
		panic(err)
	}

	return db
}

func FindAll() []model.User {
	db := initDB()
	defer db.Close()

	var users []model.User
	// ポインタで呼ぶ
	db.Find(&users)
	return users
}

func Find(id int) model.User {
	db := initDB()
	defer db.Close()

	var user model.User
	// ポインタで呼ぶ
	db.First(&user, id)
	return user
}

func Insert(name string, email string) model.User {
	db := initDB()
	defer db.Close()

	user := model.User{Name: name, Email: email}
	db.NewRecord(user)
	// ポインタで呼ぶ
	db.Create(&user)
	return user
}

func Update(id int, name string, email string) model.User {
	db := initDB()
	defer db.Close()

	var user model.User
	db.First(&user, id)
	fmt.Println(user)
	db.Model(&user).Where("id = ?", id).Updates(model.User{Name: name, Email: email})
	return user
}

func Delete(id int) {
	db := initDB()
	defer db.Close()

	var user model.User
	db.First(&user, id)
	db.Delete(user)
}

// https://qiita.com/kai1993/items/389cdec6a01bd527525b
