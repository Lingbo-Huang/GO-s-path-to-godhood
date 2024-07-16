package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	// omitempty 表示如果该字段为零值，则在JSON中忽略
	Age   int    `json:"age,omitempty"`
	Email string `json:"email"`
	// `-` 表示该字段在JSON中忽略
	Password string `json:"-"`
}

func main() {
	p := Person{
		Name:  "Alice",
		Age:   0, // Age 为零值，使用 omitempty 会在 JSON 中忽略该字段
		Email: "hlb@example.com",
	}

	jsonData, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Error marshaling to JSON:", err)
		return
	}
	fmt.Println("JSON string:", string(jsonData))
}
