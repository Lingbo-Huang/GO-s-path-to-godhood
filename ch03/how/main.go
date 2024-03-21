package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

func printSlice(data []string) {
	data[0] = "java"
	for i := 0; i < 10; i++ {
		data = append(data, strconv.Itoa(i))
	}
}

type slice struct {
	array unsafe.Pointer //指向连续内存，存储实际数据的数组指针
	len   int            //切片中元素数量
	cap   int            //array数组长度
}

func main() {
	courses := []string{"go", "cpp", "gin"}
	printSlice(courses)
	fmt.Println(courses)

	var data []int
	for i := 0; i < 2000; i++ {
		data = append(data, i)
		fmt.Printf("len: %d, cap: %d\r\n", len(data), cap(data))
	}
}
