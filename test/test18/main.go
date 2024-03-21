package main

import "fmt"

func main() {
	// 创建一个空接口类型的字典
	myDict := make(map[string]interface{})

	// 存储不同类型的值
	myDict["name"] = "John"
	myDict["age"] = 30
	myDict["city"] = "New York"
	myDict["data"] = []int{1, 2, 3}

	// 获取并打印值
	name := myDict["name"].(string)
	age := myDict["age"].(int)
	city := myDict["city"].(string)
	data := myDict["data"].([]int)

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("City:", city)
	fmt.Println("Data:", data)
}
