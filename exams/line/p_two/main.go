package main

import (
	"fmt"
)

func maxSum(nums []int, n int) (sum int) {
	if n == 0 {
		return 0
	}
	dp := make([]int, n)
	dp[0] = nums[0]
	if n > 1 {
		dp[1] = max(nums[0], nums[1])
	}

	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	var n int
	fmt.Scan(&n)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}
	result := maxSum(nums, n)
	fmt.Println(result)
}
