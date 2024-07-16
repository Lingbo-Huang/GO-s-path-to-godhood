package three_sum

import "sort"

func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)

	for i, num := range nums {
		if num > 0 {
			return ans
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
	}
}
