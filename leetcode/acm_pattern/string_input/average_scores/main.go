package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 多组输入（不知道多少组）
/*
每门课的成绩分为A、B、C、D、F五个等级，为了计算平均绩点，规定A、B、C、D、F分别代表4分、3分、2分、1分、0分。

输入
有多组测试样例。每组输入数据占一行，由一个或多个大写字母组成，字母之间由空格分隔。

输出
每组输出结果占一行。如果输入的大写字母都在集合｛A,B,C,D,F｝中，则输出对应的平均绩点，结果保留两位小数。否则，输出“Unknown”。

样例输入
A B C D F
B F F C C A
D C E F
样例输出
2.00
1.83
Unknown
*/

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
)

func main() {
	defer writer.Flush()
	m := map[string]int{"A": 4.0, "B": 3.0, "C": 2.0, "D": 1.0, "F": 0.0}
	for {
		boo := scanner.Scan()
		if boo == false {
			break
		}
		if err := scanner.Err(); err != nil {
			break
		}
		input := scanner.Text()

		keywordFrequency := make(map[string]int)
		keywords := strings.Fields(input)
		for _, keyword := range keywords {
			if keyword != "" {
				keywordFrequency[keyword]++
			}
		}
		var sum float64 = 0
		count := 0
		flag := 1
		for key, value := range keywordFrequency {
			if _, ok := m[key]; ok {
				count += value
				sum += float64((m[key] * value))
			} else {
				flag = 0
			}
		}
		if flag == 1 {
			fmt.Fprintf(writer, "%.2f\n", sum/float64(count))
		} else {
			fmt.Fprintln(writer, "Unknown")
		}
		writer.Flush()
	}
}
