package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func cpuInfo(ctx context.Context) {
	//拿到请求的id
	fmt.Printf("traceid: %s\r\n", ctx.Value("traceid"))
	//可以记录日志，这次请求是哪个traceid打印的

	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("退出cpu监控")
			return
		default:
			time.Sleep(2 * time.Second)
			fmt.Println("cpu的信息")
		}
	}
}

func main() {
	wg.Add(1)
	ctx, _ := context.WithTimeout(context.Background(), 6*time.Second)
	//并不影响上面的超时，都会执行到
	valueCtx := context.WithValue(ctx, "traceid", "hlb123")

	go cpuInfo(valueCtx) //传递子context
	wg.Wait()
	fmt.Println("监控完成")
}
