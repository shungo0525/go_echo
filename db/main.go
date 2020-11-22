package main

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"  // $ go get -u github.com/go-sql-driver/mysql
)

type User struct {
	ID   int
	Name string
}

func main() {
	index()
	show()
	insert()
	update()
	delete()
}

func index() {
	db, err := sql.Open("mysql", "root:@/echo_test")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	fmt.Println("----index----")
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(user.ID, user.Name)
	}
}

func show() {
	fmt.Println("----show----")
	db, err := sql.Open("mysql", "root:@/echo_test")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

  var user User
	err = db.QueryRow("SELECT * FROM users WHERE id = ?", 1).Scan(&user.ID, &user.Name)
	switch {
	case err == sql.ErrNoRows:
		fmt.Println("record not found")
	case err != nil:
		panic(err.Error())
	default:
		fmt.Println(user.ID, user.Name)
	}
}

func insert() {
	fmt.Println("----insert----")
	db, err := sql.Open("mysql", "root:@/echo_test")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	stmtInsert, err := db.Prepare("INSERT INTO users(name) VALUES(?)")
	if err != nil {
		panic(err.Error())
	}
	defer stmtInsert.Close()

	result, err := stmtInsert.Exec("user")
	if err != nil {
		panic(err.Error())
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(lastInsertID)
}

func update() {
	fmt.Println("----update----")
	db, err := sql.Open("mysql", "root:@/echo_test")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	stmtUpdate, err := db.Prepare("UPDATE users SET name=? WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	defer stmtUpdate.Close()

	result, err := stmtUpdate.Exec("update-user", 4)
	if err != nil {
		panic(err.Error())
	}

	rowsAffect, err := result.RowsAffected()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(rowsAffect)
}

func delete() {
fmt.Println("----delete----")
	db, err := sql.Open("mysql", "root:@/echo_test")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	stmtDelete, err := db.Prepare("DELETE FROM users WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	defer stmtDelete.Close()

	result, err := stmtDelete.Exec(4)
	if err != nil {
		panic(err.Error())
	}

	rowsAffect, err := result.RowsAffected()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(rowsAffect)
}

// 事前にmysqlでDBを作成/データをinsert
// DROP DATABASE IF EXISTS echo_test;
// CREATE DATABASE echo_test;
// USE sample_test;

// CREATE TABLE users (id INT NOT NULL AUTO_INCREMENT PRIMARY KEY, name TEXT NOT NULL) DEFAULT CHARACTER SET=utf8;

// INSERT INTO users (name) VALUES ("user1"),("user2"),("user3");