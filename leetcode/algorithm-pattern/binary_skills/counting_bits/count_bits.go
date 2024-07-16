package counting_bits

/*
给定一个非负整数 num。对于 0 ≤ i ≤ num 范围中的每个数字 i ，
计算其二进制数中的 1 的数目并将它们作为数组返回。
*/

func countBits(num int) []int {
	res := make([]int, num+1)
	for i := 0; i <= num; i++ {
		res[i] = count1(i)
	}
	return res
}

func count1(num int) int {
	res := 0
	for num != 0 {
		num = num & (num - 1)
		res++
	}
	return res
}

// 动态规划求解
func countBitsDP(num int) []int {
	res := make([]int, num+1)
	for i := 1; i <= num; i++ {
		// 上一个缺1的元素+1即可
		// i 的二进制表示中1的个数等于 i & (i - 1) 的二进制表示中1的个数加1。
		// 因为 i & (i - 1) 是 i 去掉最低位的 1
		res[i] = res[i&(i-1)] + 1
	}
	return res
}
