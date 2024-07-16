package balanced_binary_tree

/*
给定一个二叉树，判断它是否是高度平衡的二叉树。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ResultType struct {
	height int
	valid  bool
}

func isBalanced(root *TreeNode) bool {
	return dfs(root).valid
}

func dfs(root *TreeNode) (result ResultType) {
	if root == nil {
		result.valid = true
		result.height = 0
		return
	}
	left := dfs(root.Left)
	right := dfs(root.Right)
	if left.valid && right.valid && sub(left.height, right.height) <= 1 {
		result.valid = true
	}
	result.height = Max(left.height, right.height) + 1
	return
}

func sub(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
