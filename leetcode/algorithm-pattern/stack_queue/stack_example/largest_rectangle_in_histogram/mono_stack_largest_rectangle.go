package largest_rectangle_in_histogram

/* 题目
https://leetcode-cn.com/problems/largest-rectangle-in-histogram/
给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
求在该柱状图中，能够勾勒出来的矩形的最大面积。
*/

/* 单调栈
栈中存放了 j 值。从栈底到栈顶，j 的值严格单调递增，同时对应的高度值也严格单调递增；

当我们枚举到第 i 根柱子时，我们从栈顶不断地移除 height[j]≥height[i]的 j 值。
在移除完毕后，栈顶的 j 值就一定满足 height[j]<height[i]，此时 j 就是 i 左侧且最近的小于其高度的柱子。

这里会有一种特殊情况。如果我们移除了栈中所有的 j 值，那就说明 i 左侧所有柱子的高度都大于 height[i]，
那么我们可以认为 i 左侧且最近的小于其高度的柱子在位置 j=−1 ，它是一根「虚拟」的、高度无限低的柱子。
这样的定义不会对我们的答案产生任何的影响，我们也称这根「虚拟」的柱子为「哨兵」。
我们再将 i 放入栈顶。
*/

func largestRectangleArea(heights []int) int {
	n := len(heights)
	left, right := make([]int, n), make([]int, n)
	// 从左往右遍历一遍，找到所有矩形的左边界（往左走第一个高度小于该矩形的矩形下标）
	mono_stack := []int{}
	for i := 0; i < n; i++ {
		for len(mono_stack) > 0 && heights[mono_stack[len(mono_stack)-1]] >= heights[i] {
			// 栈顶元素（矩形下标）的高度不比当前矩形的高度小就出栈
			mono_stack = mono_stack[:len(mono_stack)-1]
		}
		if len(mono_stack) == 0 {
			// 如果已经弹完了，栈内没有元素了，可以假设矩形的最左边有个下标为-1，高度为0的矩形
			left[i] = -1
		} else {
			// 当前矩形的左边界为栈顶元素(左边最近的小于当前矩形高度的矩形下标)
			left[i] = mono_stack[len(mono_stack)-1]
		}
		// 当前矩形下标入栈
		mono_stack = append(mono_stack, i)
	}
	mono_stack = []int{}
	for i := n - 1; i >= 0; i++ {
		for len(mono_stack) > 0 && heights[mono_stack[len(mono_stack)-1]] >= heights[i] {
			mono_stack = mono_stack[:len(mono_stack)-1]
		}
		if len(mono_stack) == 0 {
			// 如果已经弹完了，栈内没有元素了，可以假设矩形的最右边有个下标为n，高度为0的矩形
			right[i] = n
		} else {
			right[i] = mono_stack[len(mono_stack)-1]
		}
		mono_stack = append(mono_stack, i)
	}
	ans := 0
	for i := 0; i < n; i++ {
		ans = max(ans, (right[i]-left[i]-1)*heights[i])
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// 优化：只需要一次遍历就能求出左右边界
/*
在求i的左边界的时候，当要弹出栈顶元素j时，说明j的高度大于等于i的高度
那么反过来，i其实是j的右边最近的小于它的高度的矩形下标
因为：假设在弹出j之前弹出过x，那么x的高度一定也大于i的高度，怎么能确定x的高度和j的关系呢
显然，既然j和x都在栈里，说明j的高度一定是小于x的高度的，否则在x入栈之前，j就弹出了。
说明当前元素i是j的右边第一个小于等于j的高度的
但我们需要的是第一个小于它的，所以我们可以加个判断
但其实不加判断，不管这个也是对最终答案没影响的。
因为：如果有若干个柱子的高度都等于矩形j的高度，
那么最右侧的那根柱子是可以求出正确的右边界的，而我们没有对求出左边界的算法进行任何改动，
因此最终的答案还是可以从最右侧的与矩形高度相同的柱子求得的。
*/

func largestRectangleArea2(heights []int) int {
	n := len(heights)
	left, right := make([]int, n), make([]int, n)
	// 左边界都初始化为0（默认，当栈里没有元素的时候认为左边界是-1，专门判断一下），右边界都初始化为n
	for i := 0; i < n; i++ {
		//left[i] = -1
		right[i] = n
	}
	mono_stack := []int{}
	for i := 0; i < n; i++ {
		for len(mono_stack) > 0 && heights[mono_stack[len(mono_stack)-1]] >= heights[i] {
			right[mono_stack[len(mono_stack)-1]] = i
			mono_stack = mono_stack[:len(mono_stack)-1]
		}
		if len(mono_stack) == 0 {
			left[i] = -1
		} else {
			left[i] = mono_stack[len(mono_stack)-1]
		}
		mono_stack = append(mono_stack, i)
	}
	ans := 0
	for i := 0; i < n; i++ {
		ans = max(ans, (right[i]-left[i]-1)*heights[i])
	}
	return ans
}
