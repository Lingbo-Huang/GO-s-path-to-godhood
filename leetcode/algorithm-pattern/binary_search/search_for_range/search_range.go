package search_for_range

/*
https://www.lintcode.com/problem/search-for-a-range/description
给定一个包含 n 个整数的排序数组，找出给定目标值 target 的起始和结束位置。
如果目标值不在数组中，则返回[-1, -1]
*/

/*
核心点就是找第一个 target 的索引，和最后一个 target 的索引，
所以用两次二分搜索分别找第一次和最后一次的位置
*/

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	result := make([]int, 2)
	start := 0
	end := len(nums) - 1
	for start+1 < end {
		mid := start + (end-start)/2
		if target == nums[mid] {
			// 如果相等，应该继续向左找，就能找到第一个目标值的位置
			end = mid
		} else if target < nums[mid] {
			end = mid
		} else if target > nums[mid] {
			start = mid
		}
	}
	// 这种写法，如果start或者end位置是目标值，那么一定是第一个目标值
	if nums[start] == target {
		result[0] = start // 找第一个，先判断start处
	} else if nums[end] == target {
		result[0] = end
	} else {
		result[0] = -1
		result[1] = -1
		return result
	}
	// 找最后一个目标值位置
	start = 0
	end = len(nums) - 1
	for start+1 < end {
		mid := start + (end-start)/2
		if nums[mid] == target {
			// 如果相等，应该继续向右找，就能找到最后一个目标值的位置
			start = mid
		} else if nums[mid] > target {
			end = mid
		} else if nums[mid] < target {
			start = mid
		}
	}
	if nums[end] == target {
		result[1] = end // 找最后一个先判断end处
	} else if nums[start] == target {
		result[1] = start
	} else {
		result[0] = -1
		result[1] = -1
		return result
	}
	return result
}
