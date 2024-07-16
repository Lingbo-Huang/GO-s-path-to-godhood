package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at ", t)
		}
	}()
	time.Sleep(time.Millisecond * 3500)
	ticker.Stop()
	fmt.Println("Ticker stopped.")
}