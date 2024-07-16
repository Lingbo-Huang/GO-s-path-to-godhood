package word_break

/*
给定一个非空字符串 s 和一个包含非空单词列表的字典 wordDict，
判定 s 是否可以被空格拆分为一个或多个在字典中出现的单词。
*/

// f[i] 表示前i个字符是否可以切分，return f[len]
// f[i] = f[j] && s[j+1~i] in wordDict
func wordBreak(s string, wordDict []string) bool {
	if len(s) == 0 {
		return true
	}

	wordDictSet := make(map[string]bool)
	for _, w := range wordDict {
		wordDictSet[w] = true
	}

	f := make([]bool, len(s)+1)
	f[0] = true

	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if f[j] && wordDictSet[s[j:i]] { // 因为f[i]中i是个数，所以右开区间到i
				f[i] = true
				break
			}
		}
	}
	return f[len(s)]
}

// 用字典里单词的最大长度进行剪枝

func wordBreak2(s string, wordDict []string) bool {
	if len(s) == 0 {
		return true
	}
	f := make([]bool, len(s)+1)
	f[0] = true
	max, dict := maxLen(wordDict)
	for i := 1; i <= len(s); i++ {
		l := 0
		if i > max {
			l = i - max
		}
		for j := l; j < i; j++ {
			if f[j] && inDict(s[j:i], dict) {
				f[i] = true
				break
			}
		}
	}
	return f[len(s)]
}

func maxLen(wordDict []string) (int, map[string]bool) {
	dict := make(map[string]bool)
	max := 0
	for _, v := range wordDict {
		dict[v] = true
		if len(v) > max {
			max = len(v)
		}
	}
	return max, dict
}

func inDict(s string, dict map[string]bool) bool {
	_, ok := dict[s]
	return ok
}
