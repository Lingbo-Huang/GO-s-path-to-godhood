package main

import (
	"fmt"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	matrix := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&matrix[i])
	}
	res := 0
	for i := 0; i < n-1; i++ {
		for j := 0; j < m-1; j++ {
			if isSpecial(matrix, i, j) {
				res++
			}
		}
	}
	fmt.Println(res)
}

func isSpecial(matrix []string, row int, col int) bool {
	chars := make(map[byte]bool)
	for i := row; i <= row+1; i++ {
		for j := col; j <= col+1; j++ {
			if chars[matrix[i][j]] {
				return false
			}
			chars[matrix[i][j]] = true
		}
	}
	return true
}
