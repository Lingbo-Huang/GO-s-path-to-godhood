package moveZeroes

// 双指针，遍历两次
func moveZeroes(nums []int) {
	i, j := 0, 0
	for i != len(nums) {
		if nums[i] != 0 {
			nums[j] = nums[i]
			j++
		}
		i++
	}
	for j != len(nums) {
		nums[j] = 0
		j++
	}
}

// 模拟快速排序，以0为分界线
func moveZeroes2(nums []int) {
	i, j := 0, 0
	for i != len(nums) {
		if nums[i] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
		i++
	}
}
