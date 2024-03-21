package main

import (
	"fmt"
	"time"
)

func main() {
	var msg chan int
	msg = make(chan int, 2)
	go func(msg chan int) {
		for data := range msg {
			fmt.Println(data)
		}
		fmt.Println("all done")
	}(msg)

	msg <- 1
	msg <- 2
	close(msg)
	data2 := <-msg //close之后可以再读，但不能再写
	fmt.Println(data2)
	time.Sleep(10 * time.Second)
}
