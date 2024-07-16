package quick_sort

func QuickSort(nums []int) []int {
	quick(nums, 0, len(nums)-1)
	return nums
}

func quick(nums []int, l, r int) {
	if l >= r {
		return
	}
	index := partition(nums, l, r)
	quick(nums, l, index-1)
	quick(nums, index+1, r)
}

func partition(nums []int, l, r int) int {
	value := nums[l]
	mark := l
	for i := l + 1; i < r; i++ {
		if nums[i] < value {
			mark++
			nums[i], nums[mark] = nums[mark], nums[i]
		}
	}
	nums[l], nums[mark] = nums[mark], nums[l]
	return mark
}
