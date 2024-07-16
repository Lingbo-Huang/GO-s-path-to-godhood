package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func main() {
	hlb := Person{
		Name:  "CaiXiaobo",
		Age:   18,
		Email: "CaiXiaobo@example.com",
	}
	// json.Marshal 把结构体转换为 json
	jsonData, err := json.Marshal(hlb)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("json:", string(jsonData))
}
