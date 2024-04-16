package count_sort

/*
计数排序，平均O(n+k)，最坏O(n+k)，空间O(n+k)，稳定
要求输入的数据必须是有确定范围的整数
找出待排序数组中最大的元素，统计每个值为i的元素出现的次数，存入数组的第i项
将所有的计数累加，将每个元素i放在新数组的第C[i]项，每放入一个元素就将C[i]减1
*/

func countSort(arr []int) []int {
	maxValue := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] > maxValue {
			maxValue = arr[i]
		}
	}
	// 以数列中最大的数作为新切片的最大下标
	bucket := make([]int, maxValue)
	// 将数列中元素值为多少就将新切片的下标为多少的位置计数加一
	for i := 0; i < len(arr); i++ {
		bucket[arr[i]]++
	}
	index := 0
	// 遍历新切片，从小往大将每个计数大于0的位置的下标写入排序数组（可以往原数组里写，因为原数组已经没价值了）
	// 每写入一个，排序数组指针往后移动一下，计数数组（桶）的计数减一
	for i := 0; i < len(bucket); i++ {
		for bucket[i] > 0 {
			arr[index] = i
			index++
			bucket[i]--
		}
	}
	return arr
}
