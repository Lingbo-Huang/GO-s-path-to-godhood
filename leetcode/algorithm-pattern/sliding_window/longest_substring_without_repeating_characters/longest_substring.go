package longest_substring_without_repeating_characters

/*
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
*/

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	win := make(map[byte]int)
	left, right := 0, 0
	ans := 1
	for right < len(s) {
		c := s[right]
		right++
		win[c]++
		// 缩小窗口
		for win[c] > 1 {
			d := s[left]
			left++
			win[d]--
		}
		ans = max(right-left, ans)
	}
	return ans
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
