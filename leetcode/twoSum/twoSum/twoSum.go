package twoSum

// 官方
func TwoSum(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, x := range nums {
		if p, ok := hashTable[target-x]; ok {
			return []int{p, i}
		}
		hashTable[x] = i
	}
	return nil
}

// 最快：map的长度直接指定了nums的长度，不需要扩容
// 但实践经验告诉我们最好给1，而不是nums长度，因为会浪费内存
func TwoSum2(nums []int, target int) []int {
	hashTable := make(map[int]int, len(nums))
	for i, x := range nums {
		if p, ok := hashTable[target-x]; ok {
			return []int{p, i}
		}
		hashTable[x] = i
	}
	return []int{}
}

// 这个别人的解法是错误的，当需要返回全部匹配时用这种
func TwoSum1(nums []int, target int) []int {

	res := make([]int, 0)
	mp := make(map[int]int, 0)

	for i, v := range nums {
		temp := target - v
		j, ok := mp[temp]
		if ok {
			res = append(res, i, j)
		}
		mp[v] = i
	}

	return res
}
