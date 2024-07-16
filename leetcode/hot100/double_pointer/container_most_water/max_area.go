package container_most_water

func maxArea(height []int) int {
	l, r := 0, len(height)
	ans := 0
	for l < r {
		area := max(height[l], height[r]) * (r - l)
		ans = max(ans, area)
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return ans
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
