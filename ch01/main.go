package main

import "fmt"

var (
	name = "bobby"
	age  = 18
	ok   bool
)

func main() {
	//go是静态语言，需要先定义变量后使用,类型和赋值类型一致
	var name int
	name = 10
	age := 1
	var user1, user2, user3 = "bobby1", 1, "bobby3"
	fmt.Println(user1, user2, user3)
	fmt.Println(age)
	fmt.Println(name)
	fmt.Println(name, ok)
}
