package dfs

// 从下向上，分治
func preorderTraverse(root *TreeNode) []int {
	result := divideAndConquer(root)
	return result
}

func divideAndConquer(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	// 分治
	left := divideAndConquer(root.Left)
	right := divideAndConquer(root.Right)
	result = append(result, root.Val)
	result = append(result, left...)
	result = append(result, right...)
	return result
}
