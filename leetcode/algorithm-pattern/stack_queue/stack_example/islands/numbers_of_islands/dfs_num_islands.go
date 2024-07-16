// package numbers_of_islands
package main

import (
	"fmt"
	"strconv"
)

func numIslands(grid [][]int) int {
	var count int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				count++
				dfs(grid, i, j)
			}
		}
	}
	return count
}

func dfs(grid [][]int, r, c int) {
	if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) {
		return
	}
	if grid[r][c] == 1 {
		grid[r][c] = 2
		dfs(grid, r-1, c)
		dfs(grid, r, c+1)
		dfs(grid, r+1, c)
		dfs(grid, r, c-1)
	}
}

func main() {
	g := [][]string{
		{"1", "1", "1", "1", "0"},
		{"1", "1", "0", "1", "0"},
		{"1", "1", "0", "0", "0"},
		{"0", "0", "0", "0", "0"},
	}
	t := make([][]int, len(g))
	for i := 0; i < len(g); i++ {
		t[i] = make([]int, len(g[i]))
		for j := 0; j < len(g[i]); j++ {
			t[i][j], _ = strconv.Atoi(g[i][j])
		}
	}

	for _, k := range t {
		fmt.Println(k)
	}
	fmt.Println(numIslands(t))
}
