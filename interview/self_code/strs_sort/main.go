package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

var (
	sc = bufio.NewScanner(os.Stdin)
)

func main() {
	for sc.Scan() {
		strs := strings.Split(sc.Text(), " ")
		sort.Strings(strs)
		fmt.Println(strings.Join(strs, " "))
	}
}
