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
	jsonString := `{"name":"CaiXiaobo", "age":22, "email":"hlb@example.com"}`

	var p Person
	err := json.Unmarshal([]byte(jsonString), &p)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}
	fmt.Printf("Go Struct: %+v\n", p)
}
