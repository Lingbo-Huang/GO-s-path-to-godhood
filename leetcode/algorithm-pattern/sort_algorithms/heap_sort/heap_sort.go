package heap_sort

func HeapSort(a []int) []int {
	// 将无序数组 a 构建为一个大根堆
	// len(a)/2-1是最后一个子树的根节点
	for i := len(a)/2 - 1; i >= 0; i-- {
		sink(a, i, len(a))
	}
	// 交换 a[0] 和 a[len(a)-1] (把有序的放最后面)
	// 然后把前面这段数组继续下沉保持堆结构
	for i := len(a) - 1; i >= 1; i-- {
		// 从后往前填充有序的值
		swap(a, 0, i)
		// 前面无序的长度减1
		sink(a, 0, i)
	}
	return a
}

func sink(a []int, i int, length int) {
	for {
		// 子节点索引
		// 从 0 开始，所以子节点为 i*2+1 和 i*2+2
		l := i*2 + 1
		r := i*2 + 2
		// idx保存根、左、右三者之间较大值的索引
		idx := i
		// 存在未排序的左节点，左节点值较大，则取左节点
		if l < length && a[l] > a[idx] {
			idx = l
		}
		// 存在未排序的右节点，且值较大，取右节点
		if r < length && a[r] > a[idx] {
			idx = r
		}
		// 如果根节点较大，则不用下沉
		if idx == i {
			break
		}
		// 如果根节点较小，则交换值，并继续下沉
		swap(a, i, idx)
		// 继续下沉
		i = idx
	}
}

func swap(a []int, i, j int) {
	a[i], a[j] = a[j], a[i]
}
