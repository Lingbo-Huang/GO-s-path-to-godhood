package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age" valid:"1-100"`
}

func (p *Person) validation() bool {
	v := reflect.ValueOf(*p)            // 获取结构体对象的值，理解为真正的结构体对象
	val := v.Field(1).Interface().(int) // 获取该结构体对象的下标为1的字段值，Field返回Value类型
	// 然后用Interface()返回当前类型转任意类型，再断言int类型
	t := reflect.TypeOf(*p) // 获取结构体对象的类型，理解为：Person这个结构体定义
	// 可以等效于：t := v.Type()
	field := t.Field(1)           // 获取该结构体定义的下标为1的字段
	tag := field.Tag.Get("valid") // 获取该字段的tag部分中的key为valid的值（标签）
	fmt.Printf("tag:%v, val:%v\n", tag, val)

	result := strings.Split(tag, "-")
	var min, max int
	min, _ = strconv.Atoi(result[0])
	max, _ = strconv.Atoi(result[1])
	if val >= min && val <= max {
		return true
	}
	return false
}

func main() {
	person1 := Person{"tom", 12}
	if person1.validation() {
		fmt.Printf("%s's age is valid.\n", person1.Name)
	} else {
		fmt.Printf("%s's age is invalid.\n", person1.Name)
	}
}
