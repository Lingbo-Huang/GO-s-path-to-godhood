package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "hello"
	b := "hello"
	fmt.Println(a == b)
	name := "hello hlb"
	fmt.Println(strings.Contains(name, a))
	bytes := []rune(name)
	fmt.Println(len(bytes))
	fmt.Println(strings.Count(name, "h"))
	fmt.Println(strings.Split(name, " "))
	fmt.Println(strings.HasPrefix(name, "hello"))
	fmt.Println(strings.HasSuffix(name, "hlb"))
	fmt.Println(strings.Replace(name, "h", "hhhh", -1))
	fmt.Println(name)
	fmt.Println(strings.ToUpper(name))
	fmt.Println(strings.Trim("$$#hello #go#", "#$"))
}
