package main

import "fmt"

type Duck interface {
	Gaga()
	Walk()
	Swimming()
}

type pskDuck struct {
	legs int
}

func (pd *pskDuck) Gaga() {
	fmt.Println("gaga")
}

func (pd *pskDuck) Walk() {
	fmt.Println("wwwwww")
}

func (pd *pskDuck) Swimming() {
	fmt.Println("ssss")
}

func main() {
	var d Duck = &pskDuck{}
	d.Gaga()
}
