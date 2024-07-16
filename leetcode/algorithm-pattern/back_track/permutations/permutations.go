package permutations

/*
给定一个 没有重复 数字的序列，返回其所有可能的全排列。
*/

func permute(nums []int) [][]int {
	result := make([][]int, 0)
	list := make([]int, 0)
	// 标记这个元素是否已经添加到结果集
	visited := make([]bool, len(nums))
	backtrack(nums, visited, list, &result)
	return result
}

func backtrack(nums []int, visited []bool, list []int, result *[][]int) {
	if len(list) == len(nums) {
		ans := make([]int, len(list))
		copy(ans, list)
		*result = append(*result, ans)
		return
	}
	for i := 0; i < len(nums); i++ {
		// 已经添加过的元素，直接跳过
		if visited[i] {
			continue
		}
		// 添加元素
		list = append(list, nums[i])
		visited[i] = true
		backtrack(nums, visited, list, result)
		// 移除元素
		visited[i] = false
		list = list[0 : len(list)-1]
	}
}
