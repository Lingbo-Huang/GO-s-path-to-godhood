package search_in_rotated_sorted_array

/*
假设按照升序排序的数组在预先未知的某个点上进行了旋转。
( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。 搜索一个给定的目标值，
如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。 你可以假设数组中不存在重复的元素。
*/

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	start := 0
	end := len(nums) - 1
	for start+1 < end {
		mid := start + (end-start)/2
		if nums[mid] == target {
			return mid
		}
		// mid 落在旋转点左边
		if nums[start] < nums[mid] {
			// target 落在 start 和 mid 中间
			if nums[start] <= target && target <= nums[mid] {
				end = mid
			} else { // target 在 mid 右边
				start = mid
			}
		} else if nums[end] > nums[mid] { // mid 落在旋转点右边
			// target 在 mid 和 end 中间
			if nums[mid] <= target && target <= nums[end] {
				start = mid
			} else { // target 在 mid 左边
				end = mid
			}
		}
	}
	if nums[start] == target {
		return start
	} else if nums[end] == target {
		return end
	}
	return -1
}
