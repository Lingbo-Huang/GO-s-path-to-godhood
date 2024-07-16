package single_number3

/*
给定一个非空整数数组，除了某个元素只出现一次以外，
其余每个元素均出现了三次。找出那个只出现了一次的元素。
*/

// 求出答案的每一位：因为除了答案之外的每个元素出现三次
// 所以对应每一位1的个数都应该是三的倍数，除非答案的这一位是1
// 所以可以求出每一位的1的个数对3取余，拼凑出答案
func singleNumber(nums []int) int {
	var result int
	for i := 0; i < 64; i++ {
		sum := 0
		for j := 0; j < len(nums); j++ {
			// 低位（第 0 位）起第 i 位 1 的个数
			// 1 其实是最低位为1，高位都是0
			// &1 可以得到最低位是1还是0
			sum += (nums[j] >> i) & 1
		}
		// sum % 3 得到这一位是 1 还是 0
		// 首先能计算出该拼接的是第i位 (sum%3) << i
		// 因为每一位都是这一位是1其它是0
		// 所以拼接：异或（因为每一个数互不冲突，所以不会出现1和1在同一位）
		result ^= (sum % 3) << i
	}
	return result
}

// 哈希解法
func singleNumberHash(nums []int) int {
	freq := map[int]int{}
	for _, v := range nums {
		freq[v]++
	}
	for num, count := range freq {
		if count == 1 {
			return num
		}
	}
	return 0
}
