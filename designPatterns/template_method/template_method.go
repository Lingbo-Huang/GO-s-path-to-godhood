package main

import "fmt"

type Beverage interface {
	BoilWater() //煮开水
	Brew()      //冲泡
	PourInCup() //倒入杯中
	AddThings() //添加佐料

	WantAddThings() bool //是否加入佐料Hook
}

type template struct {
	b Beverage
}

func (t *template) MakeBeverage() {
	if t == nil {
		return
	}
	t.b.BoilWater()
	t.b.Brew()
	t.b.PourInCup()
	if t.b.WantAddThings() == true {
		t.b.AddThings()
	}
}

type MakeCoffee struct {
	template
}

func NewMakeCoffee() *MakeCoffee {
	makeCoffee := new(MakeCoffee)
	makeCoffee.b = makeCoffee
	return makeCoffee
}

func (mc *MakeCoffee) BoilWater() {
	fmt.Println("将水煮到100摄氏度")
}

func (mc *MakeCoffee) Brew() {
	fmt.Println("用水煮咖啡")
}

func (mc *MakeCoffee) PourInCup() {
	fmt.Println("将充好的咖啡倒入陶瓷杯中")
}

func (mc *MakeCoffee) AddThings() {
	fmt.Println("添加牛奶和糖")
}

func (mc *MakeCoffee) WantAddThings() bool {
	return true
}

type MakeTea struct {
	template
}

func NewMakeTea() *MakeTea {
	makeTea := new(MakeTea)
	makeTea.b = makeTea
	return makeTea
}

func (mt *MakeTea) BoilWater() {
	fmt.Println("将水煮到80摄氏度")
}

func (mt *MakeTea) Brew() {
	fmt.Println("用水冲茶叶")
}

func (mt *MakeTea) PourInCup() {
	fmt.Println("将冲好的茶水倒入茶壶中")
}

func (mt *MakeTea) AddThings() {
	fmt.Println("添加柠檬")
}

func (mt *MakeTea) WantAddThings() bool {
	return false
}

func main() {
	makeCoffee := NewMakeCoffee()
	makeCoffee.MakeBeverage()

	fmt.Println("-----------")

	mekeTea := NewMakeTea()
	mekeTea.MakeBeverage()
}
