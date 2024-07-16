package backpack

/*
给定一个背包的容量m 和一组物品，每个物品有一定的重量。
问题是要确定能够放入背包的最大重量。
*/

// f[i][j] 前i个物品， 是否能装重量j
func backPack(m int, A []int) int {
	f := make([][]bool, len(A)+1)
	for i := 0; i <= len(A); i++ {
		f[i] = make([]bool, m+1)
	}
	f[0][0] = true
	for i := 1; i <= len(A); i++ {
		for j := 0; j <= m; j++ {
			f[i][j] = f[i-1][j]
			if j >= A[i-1] && f[i-1][j-A[i-1]] {
				f[i][j] = true
			}
		}
	}
	for i := m; i >= 0; i-- {
		if f[len(A)][i] {
			return i
		}
	}
	return 0
}
