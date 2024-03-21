package main

import (
	"fmt"
	"time"
)

//var done = make(chan struct{}) //chan必须初始化

func g1(ch chan struct{}) {
	time.Sleep(1 * time.Second)
	ch <- struct{}{}
}

func g2(ch chan struct{}) {
	time.Sleep(3 * time.Second)
	ch <- struct{}{}
}

func main() {
	g1Channel := make(chan struct{})
	g2Channel := make(chan struct{})

	go g1(g1Channel)
	go g2(g2Channel)

	timer := time.NewTimer(5 * time.Second)
	for {
		select {
		case <-g1Channel:
			fmt.Println("g1 done")
		case <-g2Channel:
			fmt.Println("g2 done")
		case <-timer.C:
			fmt.Println("timeout")
			return
			//default:
			//	fmt.Println("default")
		}
	}

}
