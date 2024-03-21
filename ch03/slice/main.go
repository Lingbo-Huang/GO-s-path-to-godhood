package main

import "fmt"

func main() {
	var courses []string
	fmt.Printf("%T\r\n", courses)
	courses = append(courses, "go")
	courses = append(courses, "grpc")
	courses = append(courses, "gin")

	fmt.Println(courses[1])

	fmt.Println(courses[0:len(courses)])

	courseSlice := []string{"go", "grpc", "gin", "mysql", "elasticsearch"}
	courseSlice = append(courseSlice, "cpp", "os")
	fmt.Println(courseSlice)

	c1 := []string{"go", "grpc"}
	c2 := []string{"mysql", "es", "gin"}
	c1 = append(c1, c2[1:]...)
	fmt.Println(c1)

	//指针指向同一个地址
	c1Copy1 := c1[:]
	fmt.Println(c1Copy1)

	//复制到新的空间里，原切片改动不影响新的切片
	var c1Copy2 = make([]string, len(c1))
	copy(c1Copy2, c1)
	fmt.Println(c1Copy2)

	fmt.Println("-------------")
	c1[0] = "java"
	fmt.Println(c1Copy2)
	fmt.Println(c1Copy1)
}
