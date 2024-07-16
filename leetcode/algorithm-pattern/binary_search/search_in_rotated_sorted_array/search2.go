package search_in_rotated_sorted_array

/*
假设按照升序排序的数组在预先未知的某个点上进行了旋转。
( 例如，数组 [0,0,1,2,2,5,6] 可能变为 [2,5,6,0,0,1,2] )。
编写一个函数来判断给定的目标值是否存在于数组中。
若存在返回 true，否则返回 false。(包含重复元素)
*/

func search2(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}
	start := 0
	end := len(nums) - 1
	for start+1 < end {
		// 去重
		for start < end && nums[start] == nums[start+1] {
			start++
		}
		for start < end && nums[end] == nums[end-1] {
			end--
		}
		mid := start + (end-start)/2
		if nums[mid] == target {
			return true
		}
		// mid 在旋转点左边
		if nums[start] < nums[mid] {
			// target 在 start 和 mid 中间
			if nums[start] <= target && target <= nums[mid] {
				end = mid
			} else { // target 在 mid 右边
				start = mid
			}
		} else if nums[end] > nums[mid] {
			if nums[mid] <= target && target <= nums[end] {
				start = mid
			} else {
				end = mid
			}
		}
	}
	if nums[start] == target || nums[end] == target {
		return true
	}
	return false
}
