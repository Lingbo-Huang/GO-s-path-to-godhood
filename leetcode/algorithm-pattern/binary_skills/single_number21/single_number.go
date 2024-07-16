package single_number21

/*
给定一个整数数组 nums，其中恰好有两个元素只出现一次，
其余所有元素均出现两次。 找出只出现一次的那两个元素。
*/

func singleNumber(nums []int) (int, int) {
	// 对所有元素进行异或
	// 因为有两个只出现一次的不同的元素
	// 所以这个结果一定不为0
	xor := 0
	for _, num := range nums {
		xor ^= num
	}
	// 拿这个结果和自己的负数相与，得到这个结果的从低到高位第一个1
	// 表示这两个数字在这一位上不同
	mask := xor & (-xor)
	// 通过和这个mask相与，为0或者为1为分界线，把这个数组分成两个数组
	// 就转换成了两道single_number2的题
	// 也就是循环异或的最终结果就是要求的数
	ans := make([]int, 2)
	for _, num := range nums {
		if num&mask == 0 {
			ans[0] ^= num
		} else {
			ans[1] ^= num
		}
	}
	return ans[0], ans[1]
}
