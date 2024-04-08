package main

import (
	"fmt"
	"sync"
)

func main() {
	var ch1, ch2 = make(chan struct{}), make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(2)
	go func(s string) {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			<-ch1
			fmt.Print(s)
			ch2 <- struct{}{}
		}
		<-ch1
	}("1")
	go func(s string) {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			<-ch2
			fmt.Print(s)
			ch1 <- struct{}{}
		}
		<-ch2
	}("2")
	ch1 <- struct{}{}
	wg.Wait()
}
