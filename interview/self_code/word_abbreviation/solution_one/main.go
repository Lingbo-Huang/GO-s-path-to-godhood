package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	sc       = bufio.NewScanner(os.Stdin)
	writer   = bufio.NewWriter(os.Stdout)
	bf       = make([]byte, 20000*1024)
	readLine = func() (res []string) {
		sc.Scan()
		res = strings.Split(sc.Text(), " ")
		return
	}
)

func main() {
	sc.Buffer(bf, len(bf))
	defer writer.Flush()

	var n int
	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		words := readLine()
		abbreviation := make([]byte, 0)
		for _, word := range words {
			if len(word) > 0 {
				if word[0] >= 'a' && word[0] <= 'z' {
					abbreviation = append(abbreviation, word[0]-32)
				} else {
					abbreviation = append(abbreviation, word[0])
				}
			}
		}
		fmt.Fprintln(writer, string(abbreviation))
	}
}
