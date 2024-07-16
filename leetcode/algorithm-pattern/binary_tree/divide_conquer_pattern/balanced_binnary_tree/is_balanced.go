package balanced_binnary_tree

/*
给定一个二叉树，判断它是否是高度平衡的二叉树。
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 左边平衡 && 右边平衡 && 左右高度差 <= 1
// 用-1表示不平衡，>0表示树高

func isBalanced(root *TreeNode) bool {
	if maxDepth(root) == -1 {
		return false
	}
	return true
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	if left == -1 || right == -1 || left-right > 1 || right-left > 1 {
		return -1
	}
	if left > right {
		return left + 1
	} else {
		return right + 1
	}
}
