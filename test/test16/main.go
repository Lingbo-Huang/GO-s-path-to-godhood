package main

import (
	"encoding/json"
	"fmt"
)

// Person 结构体定义
type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	City  string `json:"city"`
	Email string `json:"email"`
}

func main() {
	// 创建一个 Person 结构体实例
	person := Person{
		Name:  "John",
		Age:   30,
		City:  "New York",
		Email: "john@example.com",
	}

	// 使用 json.Marshal 将结构体转换为 JSON 格式的数据
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 打印 JSON 格式的数据
	fmt.Println(string(jsonData))
}
