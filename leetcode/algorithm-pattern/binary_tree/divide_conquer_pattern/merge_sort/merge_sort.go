package merge_sort

func MergeSort(nums []int) []int {
	return mergesort(nums)
}

func mergesort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	mid := len(nums) / 2
	left := mergesort(nums[:mid])
	right := mergesort(nums[mid:])
	result := merge(left, right)
	return result
}

func merge(left []int, right []int) (result []int) {
	l, r := 0, 0
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}
	if l < len(left) {
		result = append(result, left[l:]...)
	}
	if r < len(right) {
		result = append(result, right[r:]...)
	}
	return result
}
