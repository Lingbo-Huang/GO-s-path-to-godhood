package main

import "fmt"

type People interface {
	Speak(string) string
}

type Student struct {
	name string
	age  int
}

func (stu *Student) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}

func main() {
	var peo People = &Student{name: "hlb", age: 18}
	fmt.Println(peo.Speak("morning"))
}

/*
使用指针类型的接收者，一方面因为和new返回指针相匹配
另一方面，避免每次调用方法都创建对象副本，而且使用指针接收器可以修改结构体的字段
*/
