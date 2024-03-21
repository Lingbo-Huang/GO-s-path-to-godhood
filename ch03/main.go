package main

import "fmt"

func main() {
	var courses1 [3]string
	var courses2 [2]string
	fmt.Printf("%T\r\n", courses1)
	fmt.Printf("%T\r\n", courses2)

	courses1[0] = "go"
	courses1[1] = "grpc"
	courses1[2] = "gin"

	fmt.Println(courses1)

	for _, v := range courses1 {
		fmt.Println(v)
	}

	courses3 := [3]string{"hlb", "love", "go"}
	for _, v := range courses3 {
		fmt.Println(v)
	}

	course4 := [3]string{2: "hh"}
	fmt.Println(course4)

	course5 := [...]string{"go", "cpp"}
	fmt.Println(course5)
}
