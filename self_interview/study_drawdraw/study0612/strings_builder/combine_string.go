package main

import (
	"fmt"
	"strings"
)

func main() {
	username := "hlb"
	var builder strings.Builder
	builder.WriteString("用户名:")
	builder.WriteString(username)
	re := builder.String()
	fmt.Println(re)

	var x interface{}
	switch i := x.(type) {
	case nil:
		fmt.Printf("x 的类型: %T", i)
	case float64:
		fmt.Printf("x 是 float64 类型")
	default:
		fmt.Printf("未知哦")
	}

}
