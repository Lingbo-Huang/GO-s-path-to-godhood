package main

import (
	"errors"
	"fmt"
)

func A() (int, error) {
	return 0, errors.New("this is an error")
}

func main() {
	if _, err := A(); err != nil {
		fmt.Println(err)
	}

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered if A: ", r)
		}
	}()
	var names map[string]string
	names["hh"] = "hhhhhhh"
	fmt.Println(names)
}
