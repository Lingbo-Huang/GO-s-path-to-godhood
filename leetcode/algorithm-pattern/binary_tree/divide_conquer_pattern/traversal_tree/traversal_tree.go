package traversal_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	result := divideAndConquer(root)
	return result
}

func divideAndConquer(root *TreeNode) []int {
	result := make([]int, 0)
	// 返回条件
	if root == nil {
		return result
	}
	// 分治
	left := divideAndConquer(root.Left)
	right := divideAndConquer(root.Right)
	// 合并
	result = append(result, root.Val)
	result = append(result, left...)
	result = append(result, right...)
	return result
}
