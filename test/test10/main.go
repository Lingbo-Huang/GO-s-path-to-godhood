package main

import (
	"fmt"
	"strings"
)

func main() {
	originalString := "Hello, World!"

	result := processString(originalString).
		toLowerCase().
		removeSpaces().
		reverseString().
		getResult()

	fmt.Println("Original String:", originalString)
	fmt.Println("Processed String:", result)
}

// stringProcessor 用于处理字符串的结构体
type stringProcessor struct {
	value string
}

// processString 创建一个新的 stringProcessor 实例
func processString(s string) *stringProcessor {
	return &stringProcessor{value: s}
}

// toLowerCase 将字符串转换为小写
func (sp *stringProcessor) toLowerCase() *stringProcessor {
	sp.value = strings.ToLower(sp.value)
	return sp
}

// removeSpaces 移除字符串中的空格
func (sp *stringProcessor) removeSpaces() *stringProcessor {
	sp.value = strings.ReplaceAll(sp.value, " ", "")
	return sp
}

////反转字符串
//func (sp *stringProcessor)reverseString1() *stringProcessor {
//	runes := []rune(sp.value)
//	for i, j := 0, len(runes) - 1; i < j; i, j = i + 1, j - 1 {
//		runes[i], runes[j] = runes[j], runes[i]
//	}
//	sp.value = string(runes)
//	return sp
//}

// reverseString 反转字符串
func (sp *stringProcessor) reverseString() *stringProcessor {
	runes := []rune(sp.value)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	sp.value = string(runes)
	return sp
}

// getResult 获取处理后的字符串
func (sp *stringProcessor) getResult() string {
	return sp.value
}
