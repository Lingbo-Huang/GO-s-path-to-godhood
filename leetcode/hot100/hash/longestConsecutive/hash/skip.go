package hash

func longestConsecutive(nums []int) int {
	// 用map存下每个值（去重，bool类型作为值, 方便后面判断是否存在）
	mp := make(map[int]bool)
	for _, num := range nums {
		mp[num] = true
	}

	ans := 0
	// 记录每个数从它开始到连续序列的结尾的长度
	// 如果前驱已经存在，那就不用算了，因为前驱那里已经算过了这个序列了，而且比当前长
	// 一直向后判断map里面是否存在，长度++，就能计算序列长度了
	for key := range mp {
		if !mp[key-1] {
			currentNum := key
			currentStreak := 1
			for mp[currentNum+1] {
				currentNum++
				currentStreak++
			}
			if ans < currentStreak {
				ans = currentStreak
			}
		}
	}
	return ans
}
