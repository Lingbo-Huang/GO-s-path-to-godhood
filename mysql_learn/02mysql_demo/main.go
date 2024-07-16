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

// 查询单条记录
func queryOne(id int) {
	var u1 user
	sqlStr := `select id, name, age from user where id=?`
	rowObj := db.QueryRow(sqlStr, id)
	rowObj.Scan(&u1.id, &u1.name, &u1.age) // Scan里面会把连接归还到连接池
	fmt.Printf("u1:%#v\n", u1)
}

// 查询多条记录
func queryMulti(id int) {
	var u user
	sqlStr := "select id, name, age from user where id > ?"
	rows, err := db.Query(sqlStr, id)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	// 关闭rows释放持有的数据库连接回连接池
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
	}
}

// 插入数据
func insert(name string, age int) {
	sqlStr := `insert into user(name, age) values(?,?)`
	ret, err := db.Exec(sqlStr, name, age)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	li_id, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d.\n", li_id)
}

// 更新数据
func update(id int, age int) {
	sqlStr := `update user set age = ? where id = ?`
	ret, err := db.Exec(sqlStr, age, id)
	if err != nil {
		fmt.Printf("updata failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Println("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("update success, affected rows:%d\n", n)
}

// 删除数据
func delete(id int) {
	sqlStr := `delete from user where id = ?`
	ret, err := db.Exec(sqlStr, id)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("delete success, affected rows:%d\n", n)
}

// 预处理方式查询多条数据
func prepareQuery() {
	sqlStr := `select id, name, age from user where id > ?`
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("id:%d,\t name:%s,\t age:%d\n", u.id, u.name, u.age)
	}
}

// 预处理方式插入多条数据
func prepareInsert() {
	sqlStr := `insert into user(name, age) values(?, ?)`
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer stmt.Close()
	var m = map[string]int{
		"张三": 12,
		"李四": 24,
		"王五": 36,
	}
	var sum int64
	for k, v := range m {
		ret, err := stmt.Exec(k, v)
		if err != nil {
			fmt.Printf("insert (name:%s, age:%d) failed, err:%v\n", k, v, err)
			return
		}
		n, err := ret.RowsAffected()
		if err != nil {
			fmt.Println("get RowsAffected failed, err:%v\n", err)
			return
		}
		sum += n
	}
	fmt.Printf("prepareInsert success, affected %d rows.\n", sum)
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("initDB failed, err:%v\n", err)
		return
	}
	fmt.Println("连接数据库成功")

	//queryOne(2)
	//queryMulti(0)
	//insert("帅帅", 18)
	//update(2, 99)
	//delete(2)
	//prepareInsert()
	prepareQuery()
}
