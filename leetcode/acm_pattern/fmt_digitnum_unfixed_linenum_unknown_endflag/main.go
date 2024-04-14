package main

import "fmt"

//用fmt,每行数字数量不固定，但知道数量；不知道行数;有结束标志

func main() {
	var t, sum int
	for {
		sum = 0
		fmt.Scan(&t)
		if t == 0 {
			break
		} else {
			a := make([]int, t)
			for i := 0; i < t; i++ {
				fmt.Scan(&a[i])
				sum += a[i]
			}
			fmt.Println(sum)
		}
	}
}
