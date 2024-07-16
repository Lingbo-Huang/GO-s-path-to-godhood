package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func initDB() (err error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/sql_test"
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	return
}

type user struct {
	ID   int
	NAME string
	AGE  int
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed, err:%v\n", err)
		return
	}
	sqlStr1 := `select id, name, age from user where id = 1`
	var u user
	err = db.Get(&u, sqlStr1)
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
	fmt.Printf("u:%#v\n", u)

	var userList []user
	sqlStr2 := `select id, name, age from user where id > 0`
	err = db.Select(&userList, sqlStr2)
	if err != nil {
		fmt.Printf("select failed, err:%v\n", err)
	}
	fmt.Printf("userList:%#v\n", userList)
}
