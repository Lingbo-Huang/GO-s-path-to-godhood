package single_number

/*
有序数组，有2N+1个数，其中N个数成对出现，仅1个数单独出现，找出那个单独出现的数.,时间复杂度为O(n)
利用异或^运算，相同数字异或结果为0，任何数与0异或等于它本身。
因此，对于成对出现的数字进行异或操作，最终结果将为0，而单独出现的数字与0进行异或操作将得到它本身
*/

func singleNumber(nums []int) int {
	result := 0
	for _, num := range nums {
		result ^= num
	}
	return result
}
