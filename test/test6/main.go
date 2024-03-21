package main

import (
	"fmt"
	"slices"
)

func main() {
	// 创建整数切片
	numbers := []int{5, 2, 9, 1, 5, 6}

	// 排序切片
	//sort.Ints(numbers)
	slices.Sort(numbers)

	// 输出带缩进的排序结果
	fmt.Println("Sorted Numbers:")
	printIndented(numbers, 2)
}

// printIndented 输出带缩进的切片元素
func printIndented(slice []int, indentSize int) {
	indent := ""
	for _, num := range slice {
		fmt.Printf("%s%d\n", indent, num)
		indent += getIndent(indentSize)
	}
}

// getIndent 生成缩进字符串
func getIndent(indentSize int) string {
	indent := ""
	for i := 0; i < indentSize; i++ {
		indent += " "
	}
	return indent
}
