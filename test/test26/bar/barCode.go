package main

import (
	"fmt"
	"learngo/test/test26/foo"
)

func main() {
	if err := foo.Open("foo"); err != nil {
		if foo.IsNotFoundError(err) {
			fmt.Println(err)
		} else {
			panic("unknow error")
		}
	}
}
