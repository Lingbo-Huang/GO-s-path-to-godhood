package main

import "fmt"

//用fmt,每行数字数量固定，且知道行数

func main() {
	var t, a, b int
	fmt.Scanln(&t)
	for i := 0; i < t; i++ {
		fmt.Scanln(&a, &b)
		fmt.Println(a + b)
	}
}
