package main

import (
	"fmt"
	"learngo/leetcode/twoSum/twoSum"
)

func main() {
	nums := []int{3, 3}
	target := 6
	fmt.Printf("%v\n", twoSum.TwoSum(nums, target))

	a := []int{3, 3, 2, 4}
	b := 6
	fmt.Println(twoSum.TwoSum2(a, b))
}
