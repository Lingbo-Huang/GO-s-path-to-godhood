package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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
