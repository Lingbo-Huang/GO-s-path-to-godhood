package main

import (
	"fmt"
)

type point struct {
	x, y int
}

func bfs(maze [][]byte) (int, int, int) {
	rows := len(maze)
	cols := len(maze[0])
	visited := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		visited[i] = make([]bool, cols)
	}

	var start point
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if maze[i][j] == 'k' {
				start = point{i, j}
				break
			}
		}
	}

	queue := []point{start}
	visited[start.x][start.y] = true
	minDistance := rows + cols + 10
	exitCount := 0
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	path := -1
	flag := 0
	for len(queue) > 0 {

		// fmt.Println(queue)
		size := len(queue)
		q := queue
		queue = []point{}
		path++
		for i := 0; i < size; i++ {
			cur := q[0]
			q = q[1:]
			if maze[cur.x][cur.y] == 'e' {
				exitCount++
				flag = 1
				minDistance = min(minDistance, path)
				continue
			}
			for _, dir := range directions {
				nextX, nextY := cur.x+dir[0], cur.y+dir[1]
				if nextX >= 0 && nextX < rows && nextY >= 0 && nextY < cols && !visited[nextX][nextY] && maze[nextX][nextY] != '*' {
					visited[nextX][nextY] = true
					queue = append(queue, point{nextX, nextY})
				}
			}
		}
	}
	return exitCount, minDistance, flag
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	var n, m, flag int
	fmt.Scan(&n, &m)
	maze := make([][]byte, n)
	for i := 0; i < n; i++ {
		maze[i] = make([]byte, m)
		fmt.Scanln(&maze[i])
		// fmt.Println(string(maze[i]))
	}

	exitCount, distance, flag := bfs(maze)
	if flag == 0 {
		fmt.Println(-1)
	} else {
		fmt.Println(exitCount, distance)
	}
}
