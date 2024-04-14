package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//数据库信息
	dsn := "root:root@tcp(127.0.0.1:3306)/mysql_learn_demo"
	// 连接数据库
	// Open不会校验用户名和密码
	db, err := sql.Open("mysql", dsn)
	if err != nil { // dsn格式不正确的时候，Open fail
		fmt.Printf("Open %s failed, err: %v\n", dsn, err)
		return
	}

	err = db.Ping() // 尝试连接数据库
	if err != nil {
		fmt.Printf("Ping %s failed, err:%v\n", dsn, err)
		return
	}

	fmt.Println("连接数据库成功！")
}

//func main() {
//	dsn := "root:root@tcp(127.0.0.1:3306)/mysql_learn_demo"
//	db, err := sql.Open("mysql", dsn)
//	if err != nil {
//		fmt.Printf("Open %s failed, err: %v\n", dsn, err)
//		return
//	}
//
//	err = db.Ping()
//	if err != nil {
//		fmt.Printf("Ping %s failed, err:%v\n", dsn, err)
//		return
//	}
//
//	fmt.Println("连接数据库成功")
//}
