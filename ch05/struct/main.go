package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Student struct {
	p     Person
	score float32
}

func (s Student) print() {
	fmt.Printf("name: %s, age: %d", s.p.name, s.p.age)
}

func main() {
	s := Student{
		Person{"bob", 18},
		90,
	}
	fmt.Println(s.p.age)
	s.print()
}
