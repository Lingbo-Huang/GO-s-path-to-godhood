package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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
	}
}
