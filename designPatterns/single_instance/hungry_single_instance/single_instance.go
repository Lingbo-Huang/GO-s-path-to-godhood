package main

import "fmt"

type singleton struct{}

// 饿汉式单例模式，main函数还没执行时，在编译时，对象就已创建
var instance *singleton = new(singleton)

func GetInstance() *singleton {
	return instance
}

func (s *singleton) SomeThing() {
	fmt.Println("单例的某个方法")
}

func main() {
	s1 := GetInstance()
	s1.SomeThing()

	s2 := GetInstance()

	if s1 == s2 {
		fmt.Println("s1 == s2")
	}
}
