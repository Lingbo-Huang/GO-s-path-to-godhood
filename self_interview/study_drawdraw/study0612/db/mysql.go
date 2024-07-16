package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/sql_test")
	if err != nil {
		fmt.Printf("Error connecting to database: %s\n", err)
		return
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM user")
	if err != nil {
		fmt.Printf("Error executing SQL query: %s\n", err)
		return
	}
	for rows.Next() {
		var id int
		var name string
		var age int
		err := rows.Scan(&id, &name, &age)
		if err != nil {
			fmt.Printf("Scan error: %s\n", err)
			return
		}
		fmt.Printf("id: %d, name: %s, age: %d\n", id, name, age)
	}
	defer rows.Close()
}
