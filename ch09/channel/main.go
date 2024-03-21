package main

import (
	"fmt"
	"time"
)

func main() {
	var msg chan string

	msg = make(chan string, 0)
	go func(msg chan string) {
		data := <-msg
		fmt.Println(data)
	}(msg)

	msg <- "bob"
	time.Sleep(10 * time.Second)
}
