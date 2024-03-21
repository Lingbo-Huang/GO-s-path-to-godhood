package main

import (
	"fmt"
	"strings"
)

func main() {
	//包含中文的字符串长度
	name := "imooc你好"
	bytes := []rune(name)
	fmt.Println(len(bytes))

	courseName := `"go"体系可`
	fmt.Println(courseName)

	username := "hlb"
	var builder strings.Builder
	builder.WriteString("用户名：")
	builder.WriteString(username)
	re := builder.String()
	fmt.Println(re)
}
