package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	for sc.Scan() {
		strs := strings.Split(sc.Text(), " ")
		sort.Strings(strs)
		fmt.Println(strings.Join(strs, " "))
	}
}
