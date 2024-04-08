package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//第一行输入行数，后面多行输入

/*
输出一个词组中每个单词的首字母的大写组合。

输入
输入的第一行是一个整数n，表示一共有n组测试数据。
接下来有n行，每组测试数据占一行，每行有一个词组，每个词组由一个或多个单词组成；每组的单词个数不超过10个，每个单词有一个或多个大写或小写字母组成；
单词长度不超过10，由一个或多个空格分隔这些单词。

输出
请为每组测试数据输出规定的缩写，每组输出占一行。

样例输入
1
ad dfa     fgs
样例输出
ADF
*/

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscanf(reader, "%d\n", &n)

	for i := 0; i < n; i++ {
		input, _ := reader.ReadString('\n') //读取一行输入
		words := strings.Fields(input)      //分割单词
		abbreviation := make([]byte, 0)     //存储缩写

		for _, word := range words {
			if len(word) > 0 {
				if word[0] >= 'a' && word[0] <= 'z' {
					abbreviation = append(abbreviation, word[0]-32) //把首字母小写转换为大写，并添加到缩写里
				} else {
					abbreviation = append(abbreviation, word[0])
				}
			}
		}
		fmt.Fprintln(writer, string(abbreviation)) //输出缩写
	}
	writer.Flush()
}
