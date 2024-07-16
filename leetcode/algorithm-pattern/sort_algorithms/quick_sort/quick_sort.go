package quick_sort

/*
快排
*/

func QuickSort(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

// 原地交换，所以传入交换索引
func quickSort(nums []int, start, end int) {
	if start < end {
		pivot := partition(nums, start, end)
		quickSort(nums, start, pivot-1)
		quickSort(nums, pivot+1, end)
	}
}

// 分区
func partition(nums []int, start, end int) int {
	// 选取最后一个元素作为基准pivot
	p := nums[end]
	i := start // i 始终指向小于基准的元素的下一个
	// 最后一个值就是基准所以不用比较
	for j := start; j < end; j++ {
		if nums[j] < p {
			swap(nums, i, j)
			i++
		}
	}
	swap(nums, i, end)
	return i
}

// 交换元素
func swap(nums []int, i, j int) {
	t := nums[i]
	nums[i] = nums[j]
	nums[j] = t
}
