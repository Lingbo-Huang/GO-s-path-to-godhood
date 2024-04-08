package main

import "fmt"

//用fmt，每行数字数量固定，但行数不知道

func main() {
	a := 0
	b := 0
	for {
		n, _ := fmt.Scanln(&a, &b)
		if n == 0 {
			break
		} else {
			fmt.Printf("%d\n", a+b)
		}
	}
}
