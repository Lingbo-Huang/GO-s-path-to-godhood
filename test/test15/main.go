package main

//使用匿名结构体解析json数据
import (
	"encoding/json"
	"fmt"
)

// JSONData 包含要解析的 JSON 数据的结构定义
var JSONData = []byte(`
{
  "name": "John",
  "age": 30,
  "city": "New York",
  "email": "john@example.com"
}
`)

func main() {
	// 定义匿名结构体，用于解析 JSON 数据
	var person struct {
		Name  string `json:"name"`
		Age   int    `json:"age"`
		City  string `json:"city"`
		Email string `json:"email"`
	}

	// 使用 Unmarshal 解析 JSON 数据到匿名结构体
	err := json.Unmarshal(JSONData, &person)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 打印解析得到的数据
	fmt.Println("Name:", person.Name)
	fmt.Println("Age:", person.Age)
	fmt.Println("City:", person.City)
	fmt.Println("Email:", person.Email)
}
