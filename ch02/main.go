package main

import "fmt"

func main() {
	var c byte
	c = 'a'
	c1 := 97
	var c2 rune
	c2 = '黄'
	fmt.Printf("c = %c \n", c)
	fmt.Printf("c1 = %c \n", c1)
	fmt.Printf("c2 = %c", c2)
}
