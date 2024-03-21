package main

import "fmt"

func a() (int, bool) {
	return 0, true
}

func main() {
	var _ int
	_, ok := a()
	if ok {
		fmt.Println("他给我返回了true欸")
	}
	const (
		ERR1 = iota + 1
		ERR2
		ERR25 = "ha"
		ERR3
		ERR4 = iota
		ERR5
	)
	fmt.Println(ERR1, ERR2, ERR3, ERR4)
}
