package backpack

/*
有 n 个物品和一个大小为 m 的背包.
给定数组 A 表示每个物品的大小和数组 V 表示每个物品的价值.
问最多能装入背包的总价值是多大?
*/

func backPackValue(m int, A []int, V []int) int {
	f := make([][]int, len(A)+1)
	for i := 0; i < len(A)+1; i++ {
		f[i] = make([]int, m+1)
	}
	for i := 1; i <= len(A); i++ {
		for j := 0; j <= m; j++ {
			f[i][j] = f[i-1][j]
			if j >= A[i-1] {
				f[i][j] = max(f[i-1][j], f[i-1][j-A[i-1]]+V[i-1])
			}
		}
	}
	return f[len(A)][m]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
