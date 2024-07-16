package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scanln(&s)
	newS := strings.ReplaceAll(s, "mei", "tuan")
	fmt.Println(newS)
}
