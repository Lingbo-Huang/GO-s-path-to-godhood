package main

import "fmt"

// 接口
type Component interface {
	Operation()
}

// 具体组件
type ConcreteComponent struct{}

func (c *ConcreteComponent) Operation() {
	fmt.Println("具体组件的操作")
}

// 装饰器
type Decorator struct {
	component Component
}

func (d *Decorator) Operation() {
	fmt.Println("装饰器增强前的操作")
	d.component.Operation()
	fmt.Println("装饰器增强后的操作")
}

func main() {
	component := &ConcreteComponent{}
	decorator := &Decorator{component}
	decorator.Operation()
}
