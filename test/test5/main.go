package main

/*
词频统计
*/
import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	// 读取文本文件
	file, err := os.Open("./test/test5/example.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// 初始化词频统计的映射
	wordFreq := make(map[string]int)

	// 使用正则表达式分割文本并统计词频
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		words := splitWords(line)
		updateWordFreq(wordFreq, words)
	}

	// 检查扫描过程中是否发生错误
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// 输出词频统计结果
	for word, freq := range wordFreq {
		fmt.Printf("%s: %d\n", word, freq)
	}
}

// splitWords 使用正则表达式分割文本成单词
func splitWords(text string) []string {
	// 使用正则表达式匹配单词
	reg := regexp.MustCompile(`[[:word:]]+`)
	words := reg.FindAllString(text, -1)

	// 转换为小写
	for i, word := range words {
		words[i] = strings.ToLower(word)
	}

	return words
}

// updateWordFreq 更新词频统计映射
func updateWordFreq(wordFreq map[string]int, words []string) {
	for _, word := range words {
		// 排除空字符串
		if word != "" {
			// 更新词频映射
			wordFreq[word]++
		}
	}
}
