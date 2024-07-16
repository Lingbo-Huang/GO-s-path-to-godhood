package fibonacci_number

// fibonacci数列，递归+备忘录
func fib(N int) int {
	return dfs(N)
}

var m map[int]int = make(map[int]int)

func dfs(n int) int {
	if n < 2 {
		return n
	}
	if m[n] != 0 {
		return m[n]
	}
	ans := dfs(n-2) + dfs(n-1)
	m[n] = ans
	return ans
}
