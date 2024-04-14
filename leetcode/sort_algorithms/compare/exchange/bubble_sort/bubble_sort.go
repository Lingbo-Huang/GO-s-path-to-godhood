package bubble_sort

/*
冒泡排序
平均O(n^2), 最坏O(n^2)，空间O(1)，稳定
过程：对每一对相邻的元素比较，如果第一个比第二个大，则交换
	每轮确定一个未排序中最大的放在最后
*/

func bubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ { //第i轮
		for j := 0; j < len(arr)-1-i; j++ {
			//每一轮都从第一个开始，和相邻的比较（但最后会少i个元素）
			if arr[j] > arr[j+1] {
				//比后者大才交换，所以如果相等，则相对顺序是不变的，是稳定的
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
