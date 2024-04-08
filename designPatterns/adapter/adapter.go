package main

import "fmt"

// 适配的目标
type V5 interface {
	Use5V()
}

// 被适配的角色
type V220 struct{}

func (v *V220) Use220V() {
	fmt.Println("使用220V的电压")
}

// 适配器类
type Adapter struct {
	v220 *V220
}

func (a *Adapter) Use5V() {
	fmt.Println("使用适配器进行充电...")

	a.v220.Use220V()
}

func NewAdapter(v220 *V220) *Adapter {
	return &Adapter{v220: v220}
}

// 业务类
type Phone struct {
	v V5
}

func NewPhone(v V5) *Phone {
	return &Phone{v}
}

func (p *Phone) Charge() {
	fmt.Println("Phone进行了充电...")
	p.v.Use5V()
}

func main() {
	phone := NewPhone(NewAdapter(new(V220)))
	phone.Charge()
}
