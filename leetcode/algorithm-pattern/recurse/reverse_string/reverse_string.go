package reverse_string

/*
编写一个函数，其作用是将输入的字符串反转过来。
*/
func reverseString(s []byte) {
	res := make([]byte, 0)
	reverse(s, 0, &res)
	for i := 0; i < len(s); i++ {
		s[i] = res[i]
	}
}

func reverse(s []byte, i int, res *[]byte) {
	if i == len(s) {
		return
	}
	// 利用递归，从最大的i == len(s)前一个i = len(s)-1开始把s[i]加进去
	reverse(s, i+1, res)
	*res = append(*res, s[i])
}
