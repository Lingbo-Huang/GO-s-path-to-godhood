package main

import "fmt"

func main() {
	const PI float64 = 3.1415926
	const PI2 = 3.14
	const MY_NAME = "hlb"
	const (
		UNKNOWN = 1
		FEMALE  = 2
		MALE    = 3
	)

	const (
		x int = 16
		y
		s = "abc"
		z
	)

	fmt.Println(x, y, s, z)
}
