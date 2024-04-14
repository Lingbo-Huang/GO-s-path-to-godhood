package quick_sort_non_recursive

/*
快速排序，非递归版本
使用一个队列来存储每次需要进行分区的区间，在循环中依次处理队列中的区间
左右区间入队，处理队头区间，把队头区间划分后的两个左右区间入队
*/

func quickSort(arr []int) {
	quick(arr, 0, len(arr)-1)
}

func quick(arr []int, left int, right int) {
	// 定义一个队列，队列元素是每一个待处理的区间, 每一个元素是一个元素数量为2的数组，存放区间的起止下标
	queue := make([][2]int, 0)
	// 队列第一个元素就是整个数列的起止位置，left和right
	queue = append(queue, [2]int{left, right})
	// 遍历处理队列里的每一个元素（区间），找到该区间的基准值下标位置
	// partition函数应该将数列中小于该基准值的放它左边，大于基准值的放它右边
	for len(queue) > 0 {
		l, r := queue[0][0], queue[0][1]
		queue = queue[1:]
		index := partition(arr, l, r)
		// 区间左端点下标比基准下标小，则说明基准左边还有区间范围，要加入队列
		if l < index-1 {
			queue = append(queue, [2]int{l, index - 1})
		}
		// 区间右端点下标比基准下标大，则说明基准右边还有区间范围，要加入队列
		if index+1 < r {
			queue = append(queue, [2]int{index + 1, r})
		}
	}
}

func partition(arr []int, left int, right int) int {
	baseValue := arr[left] // 基准值
	mark := left
	for i := left + 1; i <= right; i++ {
		if arr[i] < baseValue {
			mark++
			arr[mark], arr[i] = arr[i], arr[mark]
		}
	}
	arr[left], arr[mark] = arr[mark], arr[left]
	return mark
}
