package bucket_sort

import "sort"

/*
桶排序：平均O(n+k)，最坏O(n+k)，空间O(n+k)，稳定
*/

func bucketSort(arr []int, bucketSize int) {
	maxValue := arr[0]
	minValue := arr[0]
	// 找出最大最小值
	for i := 1; i < len(arr); i++ {
		if arr[i] > maxValue {
			maxValue = arr[i]
		}
		if arr[i] < minValue {
			minValue = arr[i]
		}
	}
	// 计算最大最小值的距离
	d := maxValue - minValue
	// 创建一个二维切片，相当于一个切片，每个元素存一个切片（链表一样）
	bucketList := make([][]int, bucketSize)
	// 给二维切片的每一个元素都初始化一个长度为0的切片
	for i := 0; i < bucketSize; i++ {
		bucketList[i] = make([]int, 0)
	}
	// 遍历待排序数组
	for i := 0; i < len(arr); i++ {
		// (arr[i] - minValue)/d表示该元素距离最小值占总距离（最大最小的差）的比例
		// 用这个比例乘以(bucketSize - 1) （桶的最大下标），表示应该放在下标为多少的桶里
		index := (arr[i] - minValue) * (bucketSize - 1) / d
		// 在该下标的桶的位置存放该元素
		bucketList[index] = append(bucketList[index], arr[i])
	}
	// 将每个桶内部排序
	for i := 0; i < bucketSize; i++ {
		sort.Ints(bucketList[i])
	}
	// 存放排序结果的切片
	res := make([]int, 0)
	for i := 0; i < bucketSize; i++ {
		for j := 0; j < len(bucketList[i]); j++ {
			res = append(res, bucketList[i][j])
		}
	}
	copy(arr, res)
}
