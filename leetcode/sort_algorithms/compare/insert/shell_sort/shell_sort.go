// package shell_sort
package main

import "fmt"

/*
希尔排序, 平均O(n^1.3)，最坏O(n^2),空间O(1)，不稳定
升级的插入排序，为了减少移动元素的次数，将元素按步长减小来进行插入排序（每个分组元素数量很少）
比如每次步长折半，第一次步长为总数的一半
每个分组内有序
最后一次步长为1，进行插入排序，整体有序
*/

/*
0 1 2 3 4 5 6 7
3 7 2 4 8 9 1 0
第一次 gap = 8/2 = 4
0 4; 1 5; 2 6; 3 7 四组
最外面循环是设置gap
然后循环每个分组（gap个分组）
然后对每个分组插入排序：从分组内第一个未排序元素开始，每个元素进行插入
插入过程：对已排序的数列从后往前遍历找到合适的位置
*/
func shellSort(arr []int) {
	n := len(arr)
	// 步长每次减半，直到为1
	for gap := n / 2; gap > 0; gap = gap / 2 {
		// 对每个分组都执行一次插入排序
		for i := 0; i < gap; i++ {
			// 下标0是已排序的，gap是和0同组的第一个未排序的元素下标
			for j := i + gap; j < n; j += gap {
				pos := j - gap //最后一个已排序的下标
				cur := arr[j]  //待排序的第一个元素值
				// i是已排序的第一个元素下标
				for pos >= i && cur < arr[pos] {
					arr[pos+gap] = arr[pos]
					pos = pos - gap
				}
				// 插入
				arr[pos+gap] = cur
			}
		}
	}
}

func print(arr []int) {
	for _, value := range arr {
		fmt.Printf("%d ", value)
	}
	fmt.Println()
}
func main() {
	a := []int{3, 7, 2, 4, 8, 9, 1, 0}
	print(a)
	shellSort(a)
	print(a)
}
