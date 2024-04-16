package two_path_merge

/*
归并排序，平均O(nlog(n))，最坏O(nlog(n))，空间O(n)，稳定
将长度为n的输入序列分成两个长度为n/2的子序列，对两个子序列使用归并排序
将两个排序好的子序列合并成最终的排序序列
*/

// 归并排序
func mergeSort(arr []int) []int {
	n := len(arr)
	if n < 2 {
		return arr
	}
	// 递归地归并排序左右子序列，然后合并
	return merge(mergeSort(arr[:len(arr)/2]), mergeSort(arr[len(arr)/2:]))
}

// 合并
func merge(left, right []int) []int {
	res := make([]int, 0)                 //创建一个空切片存合并结果
	for len(left) > 0 && len(right) > 0 { // 两个子序列都还有待合并元素时
		// 取两个子序列中第一个元素进行比较，将较小的那个加到结果切片里
		if left[0] <= right[0] {
			res = append(res, left[0])
			left = left[1:]
		} else {
			res = append(res, right[0])
			right = right[1:]
		}
	}
	// 结束循环后，其中某个子序列可能还有剩下没有取出的元素，需要整体添加到结果序列里
	if len(left) > 0 {
		res = append(res, left...)
	}
	if len(right) > 0 {
		res = append(res, right...)
	}
	return res
}
