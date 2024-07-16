package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	str := "hello,world"
	strBytes := []byte(str)
	strChan := make(chan byte, len(str))
	/*
		对字符串for range时，第二个返回值是字符的ASCII码，需要转成byte
		或者直接strBytes := []byte(str)然后遍历strBytes
	*/
	for _, c := range strBytes {
		strChan <- c
	}
	close(strChan) // channel用完必须close

	control := make(chan struct{})

	go func() {
		defer wg.Done()
		for {
			ball, ok := <-control
			if ok {
				pr1, ok1 := <-strChan
				if ok1 {
					fmt.Printf("go2: %c\n", pr1)
				} else {
					close(control)
					return
				}
				control <- ball
			} else {
				return
			}
		}
	}()

	go func() {
		defer wg.Done()
		for {
			time.Sleep(5 * time.Millisecond)
			ball, ok := <-control
			if ok {
				pr1, ok1 := <-strChan
				if ok1 {
					fmt.Printf("go1: %c\n", pr1)
				} else {
					close(control)
					return
				}
			} else {
				return
			}
			control <- ball
		}
	}()
	control <- struct{}{}
	wg.Wait()
}
