package main

import "fmt"

type Goods struct {
	Kind string //商品种类
	Fact bool   //商品真伪
}

// 抽象层
type Shopping interface {
	Buy(goods *Goods)
}

// 实现层
type KoreaShopping struct{}

func (ks *KoreaShopping) Buy(goods *Goods) {
	fmt.Println("去韩国进行购物，买了", goods.Kind)
}

type AmericaShopping struct{}

func (as *AmericaShopping) Buy(goods *Goods) {
	fmt.Println("去韩国进行购物，买了", goods.Kind)
}

type AfricaShopping struct{}

func (as *AfricaShopping) Buy(goods *Goods) {
	fmt.Println("去韩国进行购物，买了", goods.Kind)
}

// 海外代理
type OverSeasProxy struct {
	shopping Shopping
}

func NewProxy(shopping Shopping) Shopping {
	return &OverSeasProxy{shopping: shopping}
}

func (op *OverSeasProxy) Buy(goods *Goods) {
	// 1 先辨别真伪
	if op.distinguish(goods) == true {
		// 2 调用具体要被代理的购物方式的Buy()方法
		op.shopping.Buy(goods)
		// 3 海关安检
		op.check(goods)
	}
}

func (op *OverSeasProxy) distinguish(goods *Goods) bool {
	fmt.Println("对[", goods.Kind, "] 进行了辨别真伪。")
	if goods.Fact == false {
		fmt.Println("发现假货", goods.Kind, "，不应该购买。")
	}
	return goods.Fact
}

func (op *OverSeasProxy) check(goods *Goods) {
	fmt.Println("对[", goods.Kind, "] 进行了海关检查，成功带回祖国。")
}

// 业务逻辑层
func main() {
	g1 := Goods{
		Kind: "韩国面膜",
		Fact: true,
	}

	g2 := Goods{
		Kind: "CET4证书",
		Fact: false,
	}

	var shopping Shopping
	shopping = new(KoreaShopping)

	var proxy Shopping
	proxy = NewProxy(shopping)
	proxy.Buy(&g1)
	proxy.Buy(&g2)
}
