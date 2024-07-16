package main

import (
	"fmt"
	"strconv"
)

func countTwo(l, r int) (count int) {
	count = 0
	for i := l; i <= r; i++ {
		str := strconv.Itoa(i)
		for _, c := range str {
			if c == '2' {
				count++
			}
		}
	}
	return
}

func main() {
	var l, r int
	fmt.Scan(&l, &r)
	fmt.Println(countTwo(l, r))
}
