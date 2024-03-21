package main

import "fmt"

type Person struct {
	name    string
	age     int
	address string
	height  float32
}

func main() {
	var p Person
	p.age = 20

	fmt.Println(p.height)

	address := struct {
		province string
		city     string
		town     string
	}{
		"北京",
		"光明",
		"xxx",
	}

	fmt.Println(address)
}
