package main

import "fmt"

func add(a int, b int) (sum int, err error) {
	sum = a + b
	return sum, err
}

func kk(item ...int) (sum int) {
	for _, value := range item {
		sum += value
	}
	return sum
}

func cal(op string, items ...int) func() {
	switch op {
	case "+":
		return func() {
			fmt.Println("加法")
		}
	case "-":
		return func() {
			fmt.Println("减法")
		}
	default:
		return func() {
			fmt.Println("都不是")
		}
	}
}

func calG(op string, myFunc func(items ...int) int) int {
	switch op {
	case "+":
		return kk(1, 2, 3)
	default:
		return -1
	}
}

func autoIncrement() func() int {
	local := 0
	return func() int {
		local += 1
		return local
	}
}

func main() {
	//sum, _ := add(1, 3)
	//fmt.Println(sum)
	fmt.Println(kk(1, 2, 3))
	funcVar := add
	a := 1
	b := 38
	fmt.Println(funcVar(a, b))
	calFuncVar := cal("+")
	calFuncVar()
	sum := calG("+", kk)
	fmt.Println(sum)
	nextFunc := autoIncrement()
	for i := 0; i < 5; i++ {
		fmt.Println(nextFunc())
	}
	nextFunc = autoIncrement() //再次置零
}
