package longest_common_subsequence

/*
给定两个字符串 text1 和 text2，返回这两个字符串的最长公共子序列。
*/

// dp[i][j] a前i个和b前j个字符最长公共子序列
// dp[m+1][n+1]
//   ' a d c e
// ' 0 0 0 0 0
// a 0 1 1 1 1
// c 0 1 1 2 1

func longestCommonSubsequence(a, b string) int {
	dp := make([][]int, len(a)+1)
	for i := 0; i < len(a); i++ {
		dp[i] = make([]int, len(b)+1)
	}
	for i := 1; i <= len(a); i++ {
		for j := 1; j <= len(b); j++ {
			// 相等取左上元素+1，否则取左或上的较大值
			if a[i-1] == b[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[len(a)][len(b)]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
