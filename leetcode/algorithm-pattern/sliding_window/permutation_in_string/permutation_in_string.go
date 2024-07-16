package permutation_in_string

/*
给定两个字符串 s1 和 s2，写一个函数来判断 s2 是否包含 s1 的排列。
*/

func checkInclusion(s1, s2 string) bool {
	win := make(map[byte]int)
	need := make(map[byte]int)
	for i := 0; i < len(s1); i++ {
		need[s1[i]]++
	}
	left, right := 0, 0
	match := 0
	for right < len(s2) {
		c := s2[right]
		right++
		if need[c] != 0 {
			win[c]++
			if win[c] == need[c] {
				match++
			}
		}
		// 当窗口长度大于字符串长度，缩紧窗口
		for right-left >= len(s1) {
			// 当窗口长度和字符串匹配，并且里面每个字符数量也匹配时，满足条件
			if match == len(need) {
				return true
			}
			d := s2[left]
			left++
			if need[d] != 0 {
				if win[d] == need[d] {
					match--
				}
				win[d]--
			}
		}
	}
	return false
}
