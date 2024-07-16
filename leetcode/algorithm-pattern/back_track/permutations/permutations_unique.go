package permutations

import (
	"sort"
)

/*
给定一个可包含重复数字的序列，返回所有不重复的全排列。
*/

func permuteUnique(nums []int) [][]int {
	result := make([][]int, 0)
	list := make([]int, 0)
	visited := make([]bool, len(nums))
	sort.Ints(nums)
	backtrack2(nums, visited, list, &result)
	return result
}

func backtrack2(nums []int, visited []bool, list []int, result *[][]int) {
	if len(list) == len(nums) {
		ans := make([]int, len(list))
		copy(ans, list)
		*result = append(*result, ans)
	}
	for i := 0; i < len(nums); i++ {
		if visited[i] {
			continue
		}
		// 上一个元素和当前相同，并且没有访问过就跳过
		if i != 0 && nums[i] == nums[i-1] && !visited[i-1] {
			continue
		}
		list = append(list, nums[i])
		visited[i] = true
		backtrack2(nums, visited, list, result)
		visited[i] = false
		list = list[0 : len(list)-1]
	}
}
