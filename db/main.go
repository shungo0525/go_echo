package db

import (
	"../model"
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"  // $ go get -u github.com/go-sql-driver/mysql
)

// func main() {
// 	Index()
// 	// show(1)
// 	// insert("test-user", "email")
// 	// update(1, "user-updated", "email-updated")
// 	// delete(1)
// }


// funcの戻り値を指定
func Index() []model.User {
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

	var users []model.User

	for rows.Next() {
		var user model.User
		err := rows.Scan(&user.Id, &user.Name, &user.Email)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(user.Id, user.Name, user.Email)
		users =append(users, user)
		fmt.Println(users)
	}
	return users
}

func Show(id int) model.User {
	fmt.Println("----show----")
	fmt.Println(id)
	db, err := sql.Open("mysql", "root:@/go_echo")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

  var user model.User
	err = db.QueryRow("SELECT * FROM users WHERE id = ?", id).Scan(&user.Id, &user.Name, &user.Email)
	switch {
	case err == sql.ErrNoRows:
		fmt.Println("No Rows")
	case err != nil:
		panic(err.Error())
	default:
		fmt.Println(user.Id, user.Name, user.Email)
	}
	return user
}

func Insert(name string, email string) model.User {
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
	return user
}

func Update(id int, name string, email string) model.User {
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
	return user
}

func Delete(id int) model.User {
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
	return user
}

// 事前にmysqlでDBを作成/データをinsert
// DROP DATABASE IF EXISTS go_echo;
// CREATE DATABASE go_echo;
// USE go_echo;

// CREATE TABLE users (id INT NOT NULL AUTO_INCREMENT PRIMARY KEY, name TEXT NOT NULL) DEFAULT CHARACTER SET=utf8;

// SHOW COLUMNS FROM users;
// SELECT * FROM users;

// insert into users (name, email) values ("user1", "email1");
