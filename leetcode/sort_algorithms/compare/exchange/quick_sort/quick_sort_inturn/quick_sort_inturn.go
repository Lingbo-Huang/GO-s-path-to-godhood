package quick_sort_inturn

/*
快速排序：平均O(nlog(n))，最坏O(n^2)，空间O(log(n))，不稳定
过程：分治，挑一个元素作为基准，每次将比基准小的放基准前面，比基准大的放基准后面，相同的随便放
	一次结束，该基准就位于数列的中间位置，这个叫做分区(partition)操作
	递归进行，将两侧的子数列进行排序
*/

func quickSort(arr []int) {
	quick(arr, 0, len(arr)-1) //对arr从0到len(arr)-1排序
}

func quick(arr []int, left int, right int) {
	if left >= right {
		return
	}
	// 对数列第一次分区，返回基准最终所处下标
	index := partition(arr, left, right)
	// 递归调用quick对左右两个数列递归快速排序
	quick(arr, left, index-1)
	quick(arr, index+1, right)
}

// 分区：以最左边的元素为基准进行分区，使左边小于基准，右边大于基准
// return 基准下标
// 依次交换版本：从两边一起往中间走，确保左边小玉基准值，右边大于基准值
func partition(arr []int, left int, right int) int {
	baseValue := arr[left] //基准值设为最左边的元素值, 将该元素存在基准里，后面赋值覆盖掉就没风险了

	for left < right {
		// 从最右边往前走，确保都大于等于基准值
		for baseValue <= arr[right] && left < right { // 在循环里套循环，务必要确保满足外循环条件
			right--
		}
		// 如果小于基准值，则把小于基准值的数据存到left位置（这里原本Left元素已经存在基准值里了，不会丢）
		arr[left] = arr[right] // 执行之后，right位置的值已经是废值了，所以最后将baseValue放right位置
		// 从最左边往后走，确保都小于基准值
		for baseValue >= arr[left] && left < right {
			left++
		}
		// 如果大于基准值，则把大于基准值的数据存到right位置（这里原来right元素已经存在之前left的位置了，也不会因为覆盖丢失）
		arr[right] = arr[left]
	}
	//退出循环时，left>right
	arr[right] = baseValue
	return right

}
