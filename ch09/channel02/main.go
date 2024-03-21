package main

import (
	"fmt"
	"time"
)

func producer(out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i * i
	}
	close(out)
}

func consumer(in <-chan int) {
	for num := range in {
		fmt.Printf("num = %d\r\n", num)
	}
}

func main() {
	//var ch1 chan int
	//var ch2 chan<- float64 //只能写
	//var ch3 <-chan int     //只能读

	//c := make(chan int, 3)
	//var send chan<- int = c
	//var read <-chan int = c
	//
	//send <- 1
	//fmt.Println(<-read)

	c := make(chan int)
	go producer(c)
	go consumer(c)
	time.Sleep(10 * time.Second)
}
