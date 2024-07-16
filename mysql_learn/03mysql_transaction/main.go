package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type user struct {
	id   int
	name string
	age  int
}

func initDB() (err error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/sql_test"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return
	}

	err = db.Ping()
	if err != nil {
		return
	}

	//db.SetMaxOpenConns(10)
	//db.SetMaxIdleConns(5)
	return
}

func transactionDemo() {
	tx, err := db.Begin()
	if err != nil {
		fmt.Printf("begin failed, err:%v\n", err)
		return
	}

	sqlStr1 := `update user set age=age-2 where id = 1`
	sqlStr2 := `update user set age=age+2 where id = 2`

	result1, err := tx.Exec(sqlStr1)
	if err != nil {
		tx.Rollback()
		fmt.Printf("exec sqlStr1 failed, err:%v\n", err)
		return
	}
	n1, err := result1.RowsAffected()
	if err != nil {
		tx.Rollback()
		fmt.Printf("result1.RowsAffected() failed, er:%v\n", err)
		return
	}
	fmt.Printf("update affected %d rows.\n", n1)

	result2, err := tx.Exec(sqlStr2)
	if err != nil {
		tx.Rollback()
		fmt.Printf("exec sqlStr2 failed, err:%v\n", err)
		return
	}
	n2, err := result2.RowsAffected()
	if err != nil {
		tx.Rollback()
		fmt.Printf("result2.RowsAffected() failed, er:%v\n", err)
		return
	}
	fmt.Printf("update affected %d rows.\n", n2)

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		fmt.Println("提交出错，回滚")
		return
	}
	fmt.Println("事务提交成功")
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed, err:%v\n", err)
		return
	}
	fmt.Println("连接数据库成功！")
	transactionDemo()
}
