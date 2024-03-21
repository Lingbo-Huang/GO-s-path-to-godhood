package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var total int32
var wg sync.WaitGroup

//var lock sync.Mutex

func add() {
	defer wg.Done()
	for i := 0; i < 100000; i++ {
		atomic.AddInt32(&total, 1)
		//lock.Lock()
		//total += 1
		//lock.Unlock()
	}
}

func sub() {
	defer wg.Done()
	for i := 0; i < 100000; i++ {
		atomic.AddInt32(&total, -1)
		//lock.Lock()
		//total -= 1
		//lock.Unlock()
	}
}
func main() {
	wg.Add(2)
	go add()
	go sub()
	wg.Wait()
	fmt.Println(total)
}
