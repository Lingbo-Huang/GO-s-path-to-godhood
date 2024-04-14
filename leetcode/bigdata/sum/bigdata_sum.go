package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s1, s2 string
	for {
		n, _ := fmt.Scanln(&s1, &s2)
		if n == 0 {
			break
		} else {
			fmt.Printf("%s\n", bigNumberAdd(s1, s2))
		}
	}
}

func bigNumberAdd(s1, s2 string) string {
	// 从个位开始计算
	i, j := len(s1)-1, len(s2)-1
	// 存每个位的数字
	num1, num2 := 0, 0
	// 存每个位的进位
	carry := 0
	// 存最终结果
	s := ""
	for i >= 0 || j >= 0 {
		// 将字符串转为int
		if i >= 0 {
			num1, _ = strconv.Atoi(string(s1[i]))
		} else {
			num1 = 0
		}
		if j >= 0 {
			num2, _ = strconv.Atoi(string(s2[j]))
		} else {
			num2 = 0
		}
		// 计算该位的和：余数和进位
		sum := (num1 + num2 + carry) % 10
		s += strconv.Itoa(sum) //把int转字符串
		carry = (num1 + num2) / 10
		i--
		j--
	}
	if carry != 0 {
		s += "1"
	}
	return reverseString(s) // 求的每一位从个位开始往高位存进的字符串，所以需要逆转
}

func reverseString(s string) string {
	res := []byte(s)
	for left, right := 0, len(s)-1; left < right; left++ {
		res[left], res[right] = res[right], res[left]
		right--
	}
	return string(res)
}
