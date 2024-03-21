package main

import (
	"fmt"
	"time"
)

func main() {
	// 秒数
	seconds := int64(1638351234)
	// 将秒数转换为具体的时间
	t := time.Unix(seconds, 0)

	// 输出具体的时间
	fmt.Printf("秒数 %d 转换为时间：%s\n", seconds, t.Format("2006-01-02 15:04:05"))
	//有时间得到秒数
	unixTime := time.Date(2024, time.March, 7, 8, 1, 0, 0, time.UTC)
	fmt.Println(unixTime.Unix())
	fmt.Println(unixTime.UTC())
}
