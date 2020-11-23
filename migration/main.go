package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name    string
	Image   Image
	ImageID uint
}

type Image struct {
	gorm.Model
	Name string
	Url  string
}

// gorm.Modelと記載すると下記が追加される
// type Model struct {
//   ID        uint           `gorm:"primaryKey"`
//   CreatedAt time.Time
//   UpdatedAt time.Time
//   DeletedAt gorm.DeletedAt `gorm:"index"`
// }

func main() {
	var db *gorm.DB
	db = initDB()

	createDummyData(db)

	for _, id := range []uint{1, 2, 3} {
		var user User
		fmt.Println("-------------------------------")
		db.First(&user, id).Related(&user.Image)
		fmt.Printf("名前    : %s\n", user.Name)
		fmt.Printf("アイコン:\n")
		fmt.Printf("  名称: %s\n", user.Image.Name)
		fmt.Printf("  URL : %s\n", user.Image.Url)
	}
	fmt.Println("-------------------------------")
}

func initDB() *gorm.DB {
	// parseTime=trueはtime.Time型に変換
	db, err := gorm.Open("mysql", "root:@/?parseTime=true")
	if err != nil {
		panic("Failed to connect database")
	}
	// DB作成
	db = db.Exec("CREATE DATABASE IF NOT EXISTS echo_migrate;")
	db.Exec("USE echo_migrate;")

	// migrate
	db.AutoMigrate(&Image{})
	db.AutoMigrate(&User{})

	return db
}

func createDummyData(db *gorm.DB) {
	var members = []map[string]string{
		{"Name": "user1", "ImageFile": "user1.jpg"},
		{"Name": "user2", "ImageFile": "user2.jpg"},
		{"Name": "user3", "ImageFile": "user3.jpg"},
	}

	for _, member := range members {
		image := Image{
			Name: member["Name"] + "アイコン",
			Url:  "https://image.example.com/" + member["ImageFile"],
		}
		db.Create(&image)

		user := User{
			Name:    member["Name"],
			ImageID: image.ID,
		}
		db.Create(&user)
	}
}

// http://egawata.hatenablog.com/entry/2017/01/08/073313
