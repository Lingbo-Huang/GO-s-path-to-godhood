package minimum_path_sum

/*
给定一个包含非负整数的 m x n 网格，
请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
*/

/*
state: f[x][y] 从起点走到 x, y 的最短路径
function: f[x][y] = min(f[x-1][y], f[x][y-1]) + grid[x][y]
initial: f[0][0] = grid[0][0], f[i][0] = sum(0->i, 0), f[0][i] = sum(0, 0->i)
answer: f[n-1][m-1]
*/

func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	// 复用原来的矩阵列表
	// 初始化 f[i][0], f[0][i]
	for i := 1; i < len(grid); i++ {
		grid[i][0] = grid[i][0] + grid[i-1][0]
	}
	for j := 1; j < len(grid[0]); j++ {
		grid[0][j] = grid[0][j-1] + grid[0][j]
	}
	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[i]); j++ {
			grid[i][j] = min(grid[i][j-1], grid[i-1][j]) + grid[i][j]
		}
	}
	return grid[len(grid)-1][len(grid[0])-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
