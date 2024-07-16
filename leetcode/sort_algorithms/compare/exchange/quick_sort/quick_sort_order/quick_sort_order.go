// package quick_sort_order
package main

import "fmt"

func quickSort(arr []int) {
	quick(arr, 0, len(arr)-1)
}

func quick(arr []int, left int, right int) {
	if left >= right {
		return
	}
	index := partition(arr, left, right)
	quick(arr, left, index-1)
	quick(arr, index+1, right)
}

// 分区：顺序版本：从前往后遍历，数有多少个小于基准值的
// 遇到小于基准值的就将基准下标加一，并把该下标的值与当前下标的值交换
// 遇到大于基准值的不做处理
func partition(arr []int, left int, right int) int {
	// 记录基准值和它的下标
	baseValue := arr[left]
	mark := left
	for i := left + 1; i <= right; i++ {
		if arr[i] < baseValue {
			// mark始终都是在满足小于基准值的位置
			mark++ // 基准下标的步长为1，意味着每当不做处理时，mark下一个位置的元素不满足小于基准值的条件
			// 将mark位置的值与小于基准值的arr[i]交换
			// 确保最左边全是小于基准值的元素
			arr[mark], arr[i] = arr[i], arr[mark]
		}
	}
	// 将mark位置的值与最左边元素（基准值）交换，
	arr[left], arr[mark] = arr[mark], arr[left]
	// 返回基准值的下标，小于它的元素都在它左边
	return mark
}

func main() {
	a := []int{
		3, 5, 9, 0, 1, 2, 8,
	}
	quickSort(a)
	for _, val := range a {
		fmt.Print(" ", val)
	}
}
