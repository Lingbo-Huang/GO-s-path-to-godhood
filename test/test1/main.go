package main

/*
简介： 停止goroutine
父goroutine创建“退出channel”，往里写入
子goroutine在for循环里利用select判断是否能从“退出channel”里读取
能读取，说明父goroutine想让子goroutine噶了
*/
import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	timer := time.NewTimer(5 * time.Second)
	quit := make(chan struct{})
	go func() {
		defer wg.Done()
		for {
			select {
			case <-quit:
				fmt.Println("退出咯")
				return
			case <-timer.C:
				fmt.Println("超时")
			default:
				fmt.Println("我是default")
			}
		}
	}()
	quit <- struct{}{}
	time.Sleep(8 * time.Second)
	wg.Wait()
}
