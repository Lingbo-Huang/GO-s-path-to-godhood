package add

func Add(a, b int) int {
	if 0 == b {
		return a
	}
	// 不算进位，相加
	// 5+17:101+10001 -> 10100(用异或，没算进位)
	sum := a ^ b
	// 计算进位，同为1则进位，所以要先与再左移1位
	carry := (a & b) << 1
	// 调用自身，算sum和carry的和
	return Add(sum, carry)
}
