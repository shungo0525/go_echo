package db

import (
	"../model"
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"  // $ go get -u github.com/go-sql-driver/mysql
)

func Db() {
	// Index()
	// show(1)
	// insert("test-user", "email")
	// update(1, "user-updated", "email-updated")
	// delete(1)
}

func Index() {
	fmt.Println("----index----")
	db, err := sql.Open("mysql", "root:@/go_echo")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var user model.User
		err := rows.Scan(&user.Id, &user.Name, &user.Email)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(user.Id, user.Name, user.Email)
	}
}

func show(id int) {
	fmt.Println("----show----")
	db, err := sql.Open("mysql", "root:@/go_echo")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

  var user model.User
	err = db.QueryRow("SELECT * FROM users WHERE id = ?", id).Scan(&user.Id, &user.Name, &user.Email)
	switch {
	case err == sql.ErrNoRows:
		fmt.Println("record not found")
	case err != nil:
		panic(err.Error())
	default:
		fmt.Println(user.Id, user.Name, user.Email)
	}
}

func insert(name string, email string) {
	fmt.Println("----insert----")
	db, err := sql.Open("mysql", "root:@/go_echo")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	stmtInsert, err := db.Prepare("INSERT INTO users(name, email) VALUES(?, ?)")
	if err != nil {
		panic(err.Error())
	}
	defer stmtInsert.Close()

	result, err := stmtInsert.Exec(name, email)
	if err != nil {
		panic(err.Error())
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		panic(err.Error())
	}
	// fmt.Println(lastInsertID)

	var user model.User
	err = db.QueryRow("SELECT * FROM users WHERE id = ?", lastInsertID).Scan(&user.Id, &user.Name, &user.Email)
	fmt.Println(user.Id, user.Name, user.Email)
}

func update(id int, name string, email string) {
	fmt.Println("----update----")
	db, err := sql.Open("mysql", "root:@/go_echo")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	stmtUpdate, err := db.Prepare("UPDATE users SET name=?, email=? WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	defer stmtUpdate.Close()

	result, err := stmtUpdate.Exec(name, email, id)
	if err != nil {
		panic(err.Error())
	}

	// RowsAffected() 更新があったレコード数を取得
	rowsAffect, err := result.RowsAffected()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(rowsAffect)

	var user model.User
	err = db.QueryRow("SELECT * FROM users WHERE id = ?", id).Scan(&user.Id, &user.Name, &user.Email)
	fmt.Println(user.Id, user.Name, user.Email)
}

func delete(id int) {
	fmt.Println("----delete----")
	db, err := sql.Open("mysql", "root:@/go_echo")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	var user model.User
	err = db.QueryRow("SELECT * FROM users WHERE id = ?", id).Scan(&user.Id, &user.Name, &user.Email)

	stmtDelete, err := db.Prepare("DELETE FROM users WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	defer stmtDelete.Close()

	result, err := stmtDelete.Exec(id)
	if err != nil {
		panic(err.Error())
	}

	rowsAffect, err := result.RowsAffected()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(rowsAffect)

	fmt.Println(user.Id, user.Name, user.Email)
}

// 事前にmysqlでDBを作成/データをinsert
// DROP DATABASE IF EXISTS go_echo;
// CREATE DATABASE go_echo;
// USE go_echo;

// CREATE TABLE users (id INT NOT NULL AUTO_INCREMENT PRIMARY KEY, name TEXT NOT NULL) DEFAULT CHARACTER SET=utf8;

// SHOW COLUMNS FROM users;
// SELECT * FROM users;

// INSERT INTO users (name) VALUES ("user1"),("user2"),("user3");
