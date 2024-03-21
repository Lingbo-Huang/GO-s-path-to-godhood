package main

/*
简介： 类型开关 Type Switch
运行时检查变量类型
*/
import "fmt"

func printValue(v interface{}) {
	if value, ok := v.(string); ok {
		fmt.Printf("The value of v is: %v\r\n", value)
	} else {
		fmt.Printf("Oh no, it is not a string variable!\r\n")
	}
}

func printValue2(v interface{}) {
	switch value := v.(type) {
	case string:
		fmt.Printf("%v is a string!\r\n", value)
	case int:
		fmt.Printf("%v is a int!\r\n", value)
	default:
		fmt.Println("Who knows! What's this?!\r\n")
	}
}

func main() {
	v := 10
	printValue(v)
	fmt.Println("--------------")
	printValue2(v)
}
