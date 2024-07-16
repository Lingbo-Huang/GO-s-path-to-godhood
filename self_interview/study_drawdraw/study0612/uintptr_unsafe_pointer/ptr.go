package main

import (
	"fmt"
	"unsafe"
)

/*
unsafe.Pointer是通用指针类型，用于转换不同类型的指针
uintptr用于指针运算，GC不把uintptr当指针
*/

func main() {
	type Person struct {
		age  uint8
		name []string
	}

	var person *Person = new(Person)
	fmt.Println(person.age, person.name)
	age := unsafe.Pointer(uintptr(unsafe.Pointer(person)) + unsafe.Offsetof(person.age))
	name := unsafe.Pointer(uintptr(unsafe.Pointer(person)) + unsafe.Offsetof(person.name))
	*((*int8)(age)) = 10
	*((*[]string)(name)) = []string{"h", "l", "b"}
	fmt.Println(person.age, person.name)
}
