package and_bits

/*
给定范围 [m, n]，其中 0 <= m <= n <= 2147483647，
返回此范围内所有数字的按位与（包含 m, n 两端点）。
*/
/*
9	00001001
10	00001010
11	00001011
12	00001100
*/
// 从低位开始，含 0 的一定 & 的结果为0
// 所以只需要求二进制的公共前缀
// 也就是先不停地右移，直到m,n相等
// 再左移相同的步数，在末尾加0

func rangeBitwiseAnd(m int, n int) int {
	var count int
	for m != n {
		m >>= 1
		n >>= 1
		count++
	}
	return m << count
}
