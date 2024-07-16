package main

import "fmt"

func autoIncrement() func() int {
	local := 0
	return func() int {
		local += 1
		return local
	}
}
func main() {
	nextFunc := autoIncrement()
	for i := 0; i < 5; i++ {
		fmt.Println(nextFunc())
	}
	nextFunc = autoIncrement() //再次置零
}
