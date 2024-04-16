package simple_insert

/*
简单插入排序 平均O(n^2),最坏O(n^2),空间O(1)，稳定
认为第一个元素是已经排好序的，从已经排好序的元素序列从后向前扫描，分别与待排序的第一个元素比较
找到第一个小于等于该待排序元素的已排序元素，将该待排序元素插入其后
*/

func simpleInsertSort(arr []int) {
	// 第一个元素认为是已排序的，遍历待排序数列
	for i := 1; i < len(arr); i++ {
		pos := i - 1  // 有序数列最后一个下标
		cur := arr[i] //记录第一个待排序的元素，防止后面后移元素时覆盖丢失
		for pos >= 0 && arr[pos] > cur {
			arr[pos+1] = arr[pos]
			pos--
		}
		arr[pos+1] = cur
	}
}
