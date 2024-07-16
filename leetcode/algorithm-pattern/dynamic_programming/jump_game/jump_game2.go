package jump_game

/*
给定一个非负整数数组，你最初位于数组的第一个位置。
数组中的每个元素代表你在该位置可以跳跃的最大长度。
你的目标是使用最少的跳跃次数到达数组的最后一个位置。
*/

func jump(nums []int) int {
	// f[i] 从起点到当前位置的最小次数
	f := make([]int, len(nums))
	f[0] = 0
	for i := 1; i < len(nums); i++ {
		// 从起点到当前位置最大次数为i
		f[i] = i
		// 遍历前面的位置，如果能覆盖当前位置，则把当前最小值与从那个位置走一步比较，
		// 尝试更新
		for j := 0; j < i; j++ {
			if nums[j]+j >= i {
				f[i] = min(f[i], f[j]+1)
			}
		}
	}
	return f[len(nums)-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 法2：动态规划+贪心优化
func jump2(nums []int) int {
	n := len(nums)
	f := make([]int, n)
	f[0] = 0
	for i := 1; i < n; i++ {
		idx := 0
		// 取第一个能跳到当前位置的点即可
		// 因为idx是从小到大的，小的idx肯定到达需要的步数不大于后面的
		// 也就是跳跃次数的结果集是单调递增的，可以用贪心
		for idx < n && nums[idx]+idx < i {
			idx++
		}
		f[i] = f[idx] + 1
	}
	return f[n-1]
}

// 法3：
func jump3(nums []int) int {
	maxLength := 0
	end := 0
	ans := 0
	for i := 0; i < len(nums)-1; i++ {
		maxLength = max(maxLength, nums[i]+i)
		if i == end {
			end = maxLength
			ans++
		}
	}
	return ans
}
