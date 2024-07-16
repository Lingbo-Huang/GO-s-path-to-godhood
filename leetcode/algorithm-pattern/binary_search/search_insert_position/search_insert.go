package search_insert_position

/*
给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。
如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
*/

// 思路：找到第一个 >= target 的元素位置
func searchInsert(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	for start+1 < end {
		mid := start + (end-start)/2
		if nums[mid] == target {
			start = mid
		} else if nums[mid] > target {
			end = mid
		} else {
			start = mid
		}
	}
	if nums[start] >= target {
		return start
	} else if nums[end] >= target {
		return end
	} else if nums[end] < target {
		// 目标值比所有值都大
		return end + 1
	}
	return 0
}
