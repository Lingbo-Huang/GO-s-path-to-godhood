package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(time.Second * 2)
	<-timer.C
	fmt.Println("Timer expired")
	timer1 := time.NewTimer(time.Second)
	go func() {
		<-timer1.C
		fmt.Println("Timer1 expired")
	}()

	stop1 := timer1.Stop()
	if stop1 {
		fmt.Println("Timer1 stopped")
	}

	ch := make(chan int)
	timer2 := time.NewTimer(time.Second * 1)
	go func() {
		var x int
		for {
			select {
			case <-timer2.C:
				x++
				fmt.Printf("%d,%s\n", x, time.Now().Format("2006-01-02 12:09:23"))
				if x < 10 {
					timer2.Reset(time.Second)
				} else {
					ch <- x
				}
			}
		}
	}()
	fmt.Println(<-ch)
}
