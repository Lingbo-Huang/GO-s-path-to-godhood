package main

import "fmt"

// 抽象层
type AbstractApple interface {
	ShowApple()
}

type AbstractBanana interface {
	ShowBanana()
}

type AbstractPear interface {
	ShowPear()
}

// 抽象的工厂
type AbstractFactory interface {
	CreateApple() AbstractApple
	CreateBanana() AbstractBanana
	CreatePear() AbstractPear
}

// 实现层
// 中国产品族
type ChinaApple struct{}

func (ca *ChinaApple) ShowApple() {
	fmt.Println("China apple.")
}

type ChinaBanana struct{}

func (cb *ChinaBanana) ShowBanana() {
	fmt.Println("China banana.")
}

type ChinaPear struct{}

func (cp *ChinaPear) ShowPear() {
	fmt.Println("China pear.")
}

type ChinaFactory struct{}

func (cf *ChinaFactory) CreateApple() AbstractApple {
	var apple AbstractApple
	apple = new(ChinaApple)
	return apple
}

func (cf *ChinaFactory) CreateBanana() AbstractBanana {
	var banana AbstractBanana
	banana = new(ChinaBanana)
	return banana
}

func (cf *ChinaFactory) CreatePear() AbstractPear {
	var pear AbstractPear
	pear = new(ChinaPear)
	return pear
}

// 日本产品族
type JapanApple struct{}

func (ja *JapanApple) ShowApple() {
	fmt.Println("China apple.")
}

type JapanBanana struct{}

func (jb *JapanBanana) ShowBanana() {
	fmt.Println("Japan banana.")

}

type JapanPear struct{}

func (jp *JapanPear) ShowPear() {
	fmt.Println("Japan pear.")

}

type JapanFactory struct{}

func (jf *JapanFactory) CreateApple() AbstractApple {
	var apple AbstractApple
	apple = new(JapanApple)
	return apple
}

func (jf *JapanFactory) CreateBanana() AbstractBanana {
	var banana AbstractBanana
	banana = new(JapanBanana)
	return banana
}

func (jf *JapanFactory) CreatePear() AbstractPear {
	var pear AbstractPear
	pear = new(JapanPear)
	return pear
}

func main() {
	var cFac AbstractFactory
	cFac = new(ChinaFactory)

	var cApple AbstractApple
	cApple = cFac.CreateApple()
	cApple.ShowApple()
}
