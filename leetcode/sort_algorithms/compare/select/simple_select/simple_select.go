package simple_select

/*
简单选择排序 平均O(n^2)，最坏O(n^2)，空间O(1)，稳定
每次在未排序序列里找到最小元素放到已排序序列的最后
和冒泡排序不一样，冒泡排序每次是把元素两两比较，利用交换位置将最大的放最后
选择排序是只需要记录一下最小元素的下标即可，
*/

func selectSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}
