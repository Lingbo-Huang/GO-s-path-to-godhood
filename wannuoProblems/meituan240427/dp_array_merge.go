package main

import "fmt"

const MOD = 1000000007

func main() {
	// 读取输入
	var n, k int
	fmt.Scan(&n, &k)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	// dp[i][j]是指：在处理第i个元素时，数组（包括这个元素和它前面的元素）的最小值不小于j的合并方案的数量
	// 初始化动态规划数组
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, k+1)
	}

	// 处理第一个元素
	for j := 0; j <= k; j++ {
		if arr[0] >= j {
			dp[0][j] = 1 // 第一个元素的合并结果只与自身有关，如果不小于j，则结果为 1
		}
	}

	// 动态规划求解
	for i := 1; i < n; i++ {
		for j := 0; j <= k; j++ {
			// 不合并当前元素
			dp[i][j] = dp[i-1][j]
			// 合并当前元素
			if arr[i] < j {
				dp[i][j] = (dp[i][j] + dp[i-1][j-arr[i]]) % MOD
			}
		}
	}

	// 统计结果
	result := 0
	for j := 0; j <= k; j++ {
		result = (result + dp[n-1][j]) % MOD
	}

	// 输出结果
	fmt.Println(result)
}
