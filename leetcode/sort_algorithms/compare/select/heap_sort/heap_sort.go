package main

import "fmt"

/*
堆排序：平均O(nlog(n))，最坏O(nlog(n))，空间O(log(n))，不稳定
堆是完全二叉树：下标为 i 的结点的父结点下标为(i-1)/2, i的左孩子2i+1，右孩子2i+2
堆的插入使用上浮，删除使用下沉
以最小堆为例，父节点小于子节点
堆排序：因为堆的根节点（堆顶元素）是最小的，所以可以将堆顶元素和最后一个元素交换，则序列前面是无序的，最后是有序的
因为交换后可能新堆顶会违反堆的性质，需要调整，之后把堆顶和无序最后一个元素交换，则有序的又添加了一个新元素
*/

/*
堆排序（最后一个非叶子节点的序号是n/2-1的推理）
可以分两种情形考虑：
1.堆的最后一个非叶子节点若只有左子节点
2.堆的最后一个非叶子节点有左右两个子节点

堆的定义：
大顶堆：arr[i] >= arr[2i+1] && arr[i] >= arr[2i+2]
小顶堆：arr[i] <= arr[2i+1] && arr[i] <= arr[2i+2]

完全二叉树的性质之一是：如果节点序号为i，在它的左子节点序号为2i+1，右子节点序号为2i+2。
对于第一种情形：左子节点的序号为n-1，则n-1=2*i+1，推出i=n/2-1；
对于第二种情形：左子节点的序号为n-2，在n-2=2i+1，推出i=(n-1)/2-1；右子节点的序号为n-1，则n-1=2i+2，推出i=(n-1)/2-1；
很显然，当完全二叉树最后一个节点是其父节点的左子节点时，树的节点数为偶数；当完全二叉树最后一个节点是其父节点的右子节点时，树的节点数为奇数。
根据语法的特征，/ 符号取整，则若n为奇数时(n-1)/2-1=n/2-1。
*/

func main() {
	arr := []int{66, 33, 55, 22, 11, 99, 88, 77}
	buildHeap(arr) // 建堆
	fmt.Println(arr)
	// 插入使用上浮
	arr = append(arr, 1)
	up(arr)
	fmt.Println(arr)
	// 删除使用下沉
	value := pop(arr)
	fmt.Println(value, arr)
}

// 删除
func pop(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	res := arr[0]            // 保存要删除的堆顶元素，用于返回
	arr[0] = arr[len(arr)-1] // 将最后一个元素放到堆顶
	arr = arr[:len(arr)-1]   // 0..len(arr)-2，元素减少一
	down(arr, 0, len(arr))   // 下沉调整
	return res
}

// 构建堆
func buildHeap(arr []int) {
	// 从最后一个非叶子节点(最后一个堆)开始
	for i := len(arr)/2 - 1; i >= 0; i-- {
		down(arr, i, len(arr))
	}
}

// 下沉
func down(arr []int, parentId int, length int) {
	temp := arr[parentId]     // 父节点值
	childId := 2*parentId + 1 //左孩子
	for childId < length {    //下沉父节点的过程要一直判断到子节点不是其它堆的父节点为止
		// 小根堆下沉，首先保证childId里存的是最小的孩子
		// 有右孩子且右孩子小于左孩子
		if childId+1 < length && arr[childId+1] < arr[childId] {
			// if childId + 1 < length && arr[childId + 1] > arr[childId] // 大根堆要把最大的孩子存下来
			childId = childId + 1
		}
		if temp <= arr[childId] {
			break
		}
		// 如果父节点比最小的孩子大, 将当前三角里最小的存在父节点位置
		arr[parentId] = arr[childId]
		//往下一层走，知道把该父节点的值下沉到合适的位置
		parentId = childId
		childId = 2*childId + 1 //如果小于length，说明还有子节点
	}
	// 在该合适的位置放最初的那个待下沉的元素值
	arr[parentId] = temp
}

// 上浮
func up(arr []int) {
	childId := len(arr) - 1       // 最后一个节点下标
	last := arr[childId]          // 最后一个节点
	parentId := (childId - 1) / 2 // 父节点
	for childId > 0 && last < arr[parentId] {
		arr[childId] = arr[parentId]
		childId = parentId
		parentId = (parentId - 1) / 2
	}
	arr[childId] = last
}
