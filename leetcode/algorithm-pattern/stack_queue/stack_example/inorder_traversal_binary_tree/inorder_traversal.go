package inorder_traversal_binary_tree

/*
中序遍历二叉树，需要递归，所以用栈，先走到树的最左下，返回栈顶元素，让root = val.Right
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	stack := make([]*TreeNode, 0)
	for len(stack) > 0 || root != nil {
		// 走到最左下角
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		// 弹出元素
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, node.Val)
		// 往右试一下
		root = node.Right
	}
	return result
}
