package binary_search

/*
https://leetcode-cn.com/problems/binary-search/
给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target ，
写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。
*/

func search(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	for start+1 < end {
		mid := start + (end-start)/2
		if target == nums[mid] {
			end = mid
		} else if target > nums[mid] {
			start = mid
		} else if target < nums[mid] {
			end = mid
		}
	}
	if nums[start] == target {
		return start
	}
	if nums[end] == target {
		return end
	}
	return -1
}

// 对于不需要找第一个、最后一个位置、或者是没有重复元素 可以如下：
func search1(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	for start <= end {
		mid := start + (end-start)/2
		if target == nums[mid] {
			return mid
		} else if target < nums[mid] {
			end = mid - 1
		} else if target > nums[mid] {
			end = mid + 1
		}
	}
	return -1
}
