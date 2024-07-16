package find_all_anagrams_in_a_string

/*
给定一个字符串 s 和一个非空字符串 p，
找到 s 中所有是 p 的字母异位词的子串，返回这些子串的起始索引。
*/

func findAnagrams(s, p string) []int {
	win := make(map[byte]int)
	need := make(map[byte]int)
	for i := 0; i < len(p); i++ {
		need[p[i]]++
	}
	left, right := 0, 0
	match := 0
	ans := make([]int, 0)
	for right < len(s) {
		c := s[right]
		right++
		if need[c] != 0 {
			win[c]++
			if win[c] == need[c] {
				match++
			}
		}
		// 当窗口长度大于字符串长度，缩紧窗口
		for right-left >= len(p) {
			if right-left == len(p) && match == len(need) {
				ans = append(ans, left)
			}
			d := s[left]
			left++
			if need[d] != 0 {
				if win[d] == need[d] {
					match--
				}
				win[d]--
			}
		}
	}
	return ans
}
