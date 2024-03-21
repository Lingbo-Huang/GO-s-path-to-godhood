package main

import (
	"fmt"
	"sort"
)

func main() {
	// 创建有序的整数切片
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// 要查找的值
	target := 5

	// 使用二分查找查找目标值的索引
	index := sort.Search(len(numbers), func(i int) bool {
		return numbers[i] >= target
	})

	// 判断是否找到目标值
	if index < len(numbers) && numbers[index] == target {
		fmt.Printf("找到目标值 %d，索引为 %d\n", target, index)
	} else {
		fmt.Printf("未找到目标值 %d\n", target)
	}
}
