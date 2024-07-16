package main

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	City    string `json:"city"`
	Country string `json:"country"`
}

type Person struct {
	Name    string  `json:"Name"`
	Age     int     `json:"age"`
	Address Address `json:"address"`
}

func main() {
	jsonString := `{"name":"Alice", "age":30, "address":{"city":"HeNan", "country":"China"}}`

	var p Person
	err := json.Unmarshal([]byte(jsonString), &p)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}
	fmt.Printf("Go struct:%+v\n", p)

	jsonData, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Error marshaling to Json:", err)
		return
	}
	fmt.Println("Json string:", string(jsonData))
}
