package sub

import "learngo/leetcode/bit_operate/add"

// 减法：加上相反数
// 求相反数：各位取反再加一

func Negative(a int) int {
	return add.Add(^a, 1)
}

func Sub(a, b int) int {
	temp := Negative(b)
	return add.Add(a, temp)
}
