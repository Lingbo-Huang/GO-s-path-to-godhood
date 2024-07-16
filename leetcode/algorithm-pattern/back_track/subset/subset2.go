package subset

import "sort"

/*
给定一个可能包含重复元素的整数数组 nums，
返回该数组所有可能的子集（幂集）。
说明：解集不能包含重复的子集。
*/
func subsetWithDup(nums []int) [][]int {
	result := make([][]int, 0)
	list := make([]int, 0)
	// 先排序
	sort.Ints(nums)
	backtrack2(nums, 0, list, &result)
	return result
}

func backtrack2(nums []int, pos int, list []int, result *[][]int) {
	ans := make([]int, len(list))
	copy(ans, list)
	*result = append(*result, ans)
	for i := pos; i < len(nums); i++ {
		// 遇到重复元素，则不选择
		if i != pos && nums[i] == nums[i-1] {
			continue
		}
		list = append(list, nums[i])
		backtrack(nums, i+1, list, result)
		list = list[0 : len(list)-1]
	}
}
