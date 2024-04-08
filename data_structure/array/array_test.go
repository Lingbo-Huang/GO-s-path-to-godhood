package array

import (
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {
	/* 初始化数组 */
	var arr []int
	fmt.Println("array arr =", arr)
	nums := []int{1, 3, 2, 5, 4}
	fmt.Println("array nums =", nums)

	/* 随机访问 */
	randomNum := randomAccess(nums)
	fmt.Println("在 nums 中获取随机元素", randomNum)

	/* 长度扩展 */
	nums = extend(nums, 3)
	fmt.Println("将数组长度扩展至 8 ，得到 nums =", nums)

	/* 插入元素 */
	insert(nums, 6, 3)
	fmt.Println("在索引 3 处插入数字 6 ，得到 nums =", nums)

	/* 删除元素 */
	remove(nums, 2)
	fmt.Println("删除索引 2 处的元素，得到 nums =", nums)

	/* 遍历数组 */
	traverse(nums)

	/* 查找元素 */
	index := find(nums, 3)
	fmt.Println("在 nums 中查找元素 3 ，得到索引 =", index)

	index1 := find1(nums, 3)
	fmt.Println("在 nums 中查找元素 3 ，得到索引 =", index1)

	traverse1(nums)
}
