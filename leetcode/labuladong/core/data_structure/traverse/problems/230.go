package main

import (
	"fmt"
	"math/rand"
)

func randomAccess(nums []int) (randomNum int) {
	randomIndex := rand.Intn(len(nums))
	randomNum = nums[randomIndex]
	return
}

func main() {
	a := []int{1, 4, 6, 7, 9}
	fmt.Println(randomAccess(a))
}
