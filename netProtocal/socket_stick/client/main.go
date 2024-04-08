package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("dial failed, err:", err)
		return
	}
	defer conn.Close()
	for i := 0; i < 20; i++ {
		msg := "Hello, How are you?"
		conn.Write([]byte(msg)) //出现粘包，因为接收方不确定要传输的数据包的大小，我们可以对数据包封包和拆包（加上Encode和Decode的协议）
	}
}
