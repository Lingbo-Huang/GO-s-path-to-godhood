package main

import "fmt"

func main() {
	list := new([]int)
	*list = append(*list, 1)
	fmt.Println(*list)

	s1 := []int{1, 2, 3}
	s2 := []int{4, 5}
	s1 = append(s1, s2...)
	fmt.Println(s1)
}
