package make_large_island

/*
在二维地图上， 0 代表海洋，1代表陆地，
我们最多只能将一格 0 （海洋）变成 1 （陆地）。
进行填海之后，地图上最大的岛屿面积是多少？
*/

// 计算岛屿面积
func calculateAreas(index int, grid [][]int, i, j int) int {
	if !boundary(grid, i, j) {
		return 0
	}
	if grid[i][j] != 1 {
		return 0
	}
	grid[i][j] = index
	return 1 +
		calculateAreas(index, grid, i-1, j) +
		calculateAreas(index, grid, i, j+1) +
		calculateAreas(index, grid, i+1, j) +
		calculateAreas(index, grid, i, j-1)
}

// 获取指定位置周围的岛屿集合
func getIslands(grid [][]int, i, j int) map[int]bool {
	result := make(map[int]bool)
	if boundary(grid, i+1, j) && grid[i+1][j] != 0 {
		result[grid[i+1][j]] = true
	}
	if boundary(grid, i-1, j) && grid[i-1][j] != 0 {
		result[grid[i-1][j]] = true
	}
	if boundary(grid, i, j-1) && grid[i][j-1] != 0 {
		result[grid[i][j-1]] = true
	}
	if boundary(grid, i, j+1) && grid[i][j+1] != 0 {
		result[grid[i][j+1]] = true
	}
	return result
}

func boundary(grid [][]int, i, j int) bool {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) {
		return false
	}
	return true
}

func makeMaxIsland(grid [][]int) int {
	if grid == nil || len(grid) == 0 {
		return 1
	}
	result := 0
	m := make(map[int]int)
	index := 2

	// 遍历网络，为每个岛屿提供编号标记，计算每个岛屿的面积，存入map里
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				m[index] = calculateAreas(index, grid, i, j)
				index++
			}
		}
	}

	// 如果没有岛屿，返回1
	if len(m) == 0 {
		return 1
	}

	// 遍历网络，找到每个海洋位置周围的岛屿，计算总面积
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 0 {
				islands := getIslands(grid, i, j)
				if len(islands) == 0 {
					continue
				}
				sum := 0
				for k, bo := range islands {
					if bo {
						sum += m[k]
					}
				}
				result = max(result, sum+1)
			}
		}
	}
	if result == 0 {
		return m[2]
	}
	return result
}
