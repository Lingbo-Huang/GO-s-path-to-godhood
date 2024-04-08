package main

import (
	"fmt"
	"sync"
)

var once sync.Once

type singleton struct{}

var instance *singleton

func GetInstance() *singleton {
	once.Do(func() {
		instance = new(singleton)
	})
	return instance
}

func (s *singleton) SomeThing() {
	fmt.Println("单例的某方法")
}

func main() {
	s1 := GetInstance()
	s1.SomeThing()

	s2 := GetInstance()
	if s1 == s2 {
		fmt.Println("s1 == s2")
	}
}
