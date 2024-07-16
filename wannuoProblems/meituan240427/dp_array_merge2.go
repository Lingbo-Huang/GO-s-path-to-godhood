package main

import "fmt"

var MOD2 = int(1e9) + 7
var dp [][]int

func main() {
	// 读取输入
	var n, k int
	fmt.Scan(&n, &k)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	// 初始化动态规划数组
	dp = make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, sum(arr)+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	// 计算并输出结果
	fmt.Println(dfs(0, 0, arr, k))
}

func dfs(i, p int, arr []int, k int) int {
	if i == len(arr) {
		if p >= k {
			return 1
		}
		return 0
	}
	if dp[i][p] != -1 {
		return dp[i][p]
	}
	cnt := 0
	if p >= k {
		cnt += dfs(i+1, arr[i], arr, k) + dfs(i+1, p+arr[i], arr, k)
	} else {
		cnt += dfs(i+1, p+arr[i], arr, k)
	}
	cnt %= MOD2
	dp[i][p] = cnt
	return dp[i][p]
}

func sum(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}
