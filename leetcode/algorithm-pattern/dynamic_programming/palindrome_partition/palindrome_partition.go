package palindrome_partition

/*
给定一个字符串 s，将 s 分割成一些子串，使每个子串都是回文串。
返回符合要求的最少分割次数。
*/

/*
思路：用动态规划
f[i]： 前i个字符组成的子串需要最少多少次切割
f[i] = min{i-1, f[j] + 1}，其中j是i前面的，j+1~i是回文串
返回f[len(nums)]
*/

// IsPalindrome 判断一个字符串的下标i到j是否是回文串
func IsPalindrome(s string, i, j int) bool {
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func minCut(s string) int {
	if len(s) == 0 || len(s) == 1 {
		return 0
	}
	f := make([]int, len(s)+1)
	f[0] = -1
	f[1] = 0
	for i := 1; i <= len(s); i++ { // f[i]是前i个，所以i要从1一直到len(s)
		f[i] = i - 1
		for j := 0; j < i; j++ {
			if IsPalindrome(s, j, i-1) { // 判断回文串的函数里参数是下标，所以i-1
				f[i] = min(f[i], f[j]+1) // 这里也保证了整个字符串就是回文串的情况
			}
		}
	}
	return f[len(s)]
}
