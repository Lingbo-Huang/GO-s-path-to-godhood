package jump_game

/*
给定一个非负整数数组，你最初位于数组的第一个位置。
数组中的每个元素代表你在该位置可以跳跃的最大长度。
判断你是否能够到达最后一个位置。
*/
func canJump(nums []int) bool {
	k := 0
	for i := 0; i < len(nums); i++ {
		if i > k {
			return false
		}
		k = max(k, i+nums[i])
	}
	return true
}

// 方法2
func canJump2(nums []int) bool {
	if len(nums) == 0 {
		return true
	}
	f := make([]bool, len(nums))
	f[0] = true
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if f[j] == true && nums[j]+j >= i {
				f[i] = true
				//break
			}
		}
	}
	return f[len(nums)-1]
}
