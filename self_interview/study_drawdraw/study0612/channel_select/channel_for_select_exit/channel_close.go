package main

import "fmt"

/*
从关闭的channel读数据，如果通道已关闭，且没数据了，读取操作立即完成，返回零值
如果通道已关闭，但还有未被读取的数据，可以正常读取
如果试图向已经关闭的channel发送数据，会导致运行时panic
*/
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	done := make(chan struct{})

	go func() {
		for i := 0; i < 2; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	go func() {
		for {
			select {
			case val, ok := <-ch1:
				if !ok {
					fmt.Println("让我看看关闭后的val是：", val)
					fmt.Println("ch1 closed")
					done <- struct{}{}
					return
				}
				fmt.Println("Receive: ", val)
			case ch2 <- 2:
				fmt.Println("Sent to ch2")
			default:
				fmt.Println("Default case")
			}
		}
	}()

	if _, ok := <-done; ok {
		fmt.Println("得知channel关闭后处理...")
	}
	defer func() {
		fmt.Println("听说等会儿报panic(小声bb：得亏我在向已关闭chan写数据前注册)")
	}()
	ch1 <- 5
}
