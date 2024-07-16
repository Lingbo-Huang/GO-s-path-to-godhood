package main

import (
	"fmt"
	"time"
)

func main() {
	data := make(map[int]int, 10)
	for i := 1; i <= 10; i++ {
		data[i] = i
	}
	for index, value := range data {
		go func(index, value int) {
			fmt.Println("index:", index, ",value:", value)
		}(index, value)
	}
	time.Sleep(time.Second)
}
