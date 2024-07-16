package longest_increasing_subsequence

/*
给定一个无序的整数数组，找到其中最长上升子序列的长度。
*/

// f[i] 表示从0开始到i结尾的最长序列长度，f[0]=1

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 || len(nums) == 1 {
		return len(nums)
	}

	f := make([]int, len(nums))
	f[0] = 1
	for i := 1; i < len(nums); i++ {
		f[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				f[i] = max(f[i], f[j]+1)
			}
		}
	}
	result := f[0]
	for i := 1; i < len(nums); i++ {
		result = max(result, f[i])
	}
	return result
}

func max(a, b int) int {
	if a < b {
		return a
	}
	return b
}
