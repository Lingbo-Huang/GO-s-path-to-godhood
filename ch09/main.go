package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	//闭包可以使用所在函数体内的局部变量
	//go可以使函数变成协程goroutine
	//for循环时，用的就是i的地址处的i，而不是复制一份
	//goroutine异步执行，当某个协程执行时未必用的是当前的i的值，可能执行时i已经变了，所以很可能会出现i变量被重复使用的情况
	//for i := 0; i < 100; i++ {
	//	go func() {
	//		fmt.Println("匿名函数goroutine " + strconv.Itoa(i))
	//	}()
	//}
	//避免循环中的变量被重用的方法：1.复制一份到临时变量；2.让协程把循环变量作为参数传入，采用值传递
	//for i := 0; i < 100; i++ {
	//	tmp := i
	//	go func() {
	//		fmt.Println("匿名函数goroutine " + strconv.Itoa(tmp))
	//	}()
	//}

	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Println("匿名函数goroutine " + strconv.Itoa(i))
		}(i)
	}

	fmt.Println("main goroutine")
	time.Sleep(10 * time.Second)
}
