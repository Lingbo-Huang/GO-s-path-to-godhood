package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/mysql_learn_demo"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("Open %s failed, err: %v\n", dsn, err)
		return
	}

	err = db.Ping()
	if err != nil {
		fmt.Printf("Ping %s failed, err:%v\n", dsn, err)
		return
	}

	fmt.Println("连接数据库成功")
}
