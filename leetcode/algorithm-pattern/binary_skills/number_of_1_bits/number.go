package number_of_1_bits

/*
编写一个函数，输入是一个无符号整数，
返回其二进制表达式中数字位数为 ‘1’ 的个数（也被称为汉明重量）。
*/
func hammingWeight(num uint32) int {
	res := 0
	for num != 0 {
		// 每次移除最低位的1
		num = num & (num - 1)
		res++
	}
	return res
}
