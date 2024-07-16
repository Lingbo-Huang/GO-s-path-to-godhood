package min_window_substring

/*
给你一个字符串 S、一个字符串 T，
请在字符串 S 里面找出：包含 T 所有字母的最小子串
*/

func minWindow(s string, t string) string {
	if len(s) == 0 || len(t) == 0 || len(s) < len(t) {
		return ""
	}
	//维护两个数组，记录已有字符串指定字符的出现次数，和目标字符串指定字符的出现次数
	var need [128]int
	var have [128]int
	//将目标字符串指定字符的出现次数记录
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	//分别为左指针，右指针，最小长度(初始值为一定不可达到的长度)
	//已有字符串中目标字符串指定字符的出现总频次以及最小覆盖子串在原字符串中的起始位置
	left, right := 0, 0
	min := len(s) + 1
	count := 0
	start := 0
	for right < len(s) {
		r := s[right]
		// 该字符不被目标字符串需要,直接移动右指针
		if need[r] == 0 {
			right++
			continue
		}
		//当且仅当已有字符串目标字符出现的次数小于目标字符串字符的出现次数时，count才会+1
		//是为了后续能直接判断已有字符串是否已经包含了目标字符串的所有字符，不需要挨个比对字符出现的次数
		if have[r] < need[r] {
			count++
		}
		//已有字符串中目标字符出现的次数+1
		have[r]++
		//移动右指针
		right++
		for count == len(t) {
			//挡窗口的长度比已有的最短值小时，更改最小值，并记录起始位置
			if right-left < min {
				min = right - left
				start = left
			}
			l := s[left]
			if need[l] == 0 {
				left++
				continue
			}
			//如果左边即将要去掉的字符被目标字符串需要，且出现的频次正好等于指定频次，那么如果去掉了这个字符，
			//就不满足覆盖子串的条件，此时要破坏循环条件跳出循环，即控制目标字符串指定字符的出现总频次(count）-1
			if have[l] == need[l] {
				count--
			}
			//已有字符串中目标字符出现的次数-1
			have[l]--
			//移动左指针
			left++
		}
	}
	if min == len(s)+1 {
		return ""
	}
	return s[start : start+min]
}
