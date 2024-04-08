package main

import "fmt"

// 抽象层

// 水果类（抽象的接口）
type Fruit interface {
	Show()
}

// 工厂类（抽象的接口）
type AbstractFactory interface {
	CreateFruit() Fruit
}

// 基础模块层

type Apple struct {
	Fruit
}

func (apple *Apple) Show() {
	fmt.Println("I'm Apple")
}

type Banana struct {
	Fruit
}

func (banana *Banana) Show() {
	fmt.Println("I'm Banana")
}

type Pear struct {
	Fruit
}

func (pear *Pear) Show() {
	fmt.Println("I'm Pear")
}

// 基础的工厂模块
type AppleFactory struct {
	AbstractFactory
}

func (af *AppleFactory) CreateFruit() Fruit {
	var fruit Fruit
	fruit = new(Apple)
	return fruit
}

type BananaFactory struct {
	AbstractFactory
}

func (bf *BananaFactory) CreateFruit() Fruit {
	var fruit Fruit
	fruit = new(Banana)
	return fruit
}

type PearFactory struct {
	AbstractFactory
}

func (pf *PearFactory) CreateFruit() Fruit {
	var fruit Fruit
	fruit = new(Pear)
	return fruit
}

// 业务逻辑层
func main() {
	var appleFac AbstractFactory
	appleFac = new(AppleFactory)
	var apple Fruit
	apple = appleFac.CreateFruit()
	apple.Show()
}
