// package nearest_matrix
package main

import "fmt"

/*
https://leetcode-cn.com/problems/01-matrix/
给定一个由 0 和 1 组成的矩阵，找出每个元素到最近的 0 的距离。 两个相邻元素间的距离为 1
*/

func updateMatrix(matrix [][]int) [][]int {
	q := make([][]int, 0)
	// 把值为0的坐标都入队，值不为0的都置为-1
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				// 入队
				point := []int{i, j}
				q = append(q, point)
			} else {
				matrix[i][j] = -1
			}
		}
	}
	directions := [][]int{
		{0, 1},
		{0, -1},
		{-1, 0},
		{1, 0},
	}
	for len(q) != 0 {
		// 出队
		point := q[0]
		q = q[1:]
		for _, v := range directions {
			x := point[0] + v[0]
			y := point[1] + v[1]
			if x >= 0 && x < len(matrix) && y >= 0 && y < len(matrix[0]) && matrix[x][y] == -1 {
				matrix[x][y] = matrix[point[0]][point[1]] + 1
				// 将当前元素入队
				q = append(q, []int{x, y})
			}
		}
	}
	return matrix
}

func main() {
	mat := [][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}}
	fmt.Println(updateMatrix(mat))
	//[[0,0,0],[0,1,0],[1,2,1]]
}
