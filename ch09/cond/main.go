package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// 创建互斥锁和条件变量
	var mu sync.Mutex
	cond := sync.NewCond(&mu)

	// 计数器，用于控制goroutine数量
	count := 3

	// 等待goroutine
	for i := 0; i < count; i++ {
		go func(id int) {
			// 获取锁
			mu.Lock()
			defer mu.Unlock()

			fmt.Printf("goroutine %d: waiting...\n", id)

			// 等待通知，条件不满足时会阻塞
			cond.Wait()

			fmt.Printf("goroutine %d: received notification\n", id)
		}(i)
	}

	// 模拟一些耗时操作
	time.Sleep(2 * time.Second)

	// 发送通知给等待的goroutine
	fmt.Println("sending notification...")
	cond.Broadcast()

	// 等待一段时间，观察效果
	time.Sleep(1 * time.Second)
}
