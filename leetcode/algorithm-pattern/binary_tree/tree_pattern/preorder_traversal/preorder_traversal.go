package preorder_traversal

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归
func preorderTraverse(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	preorderTraverse(root.Left)
	preorderTraverse(root.Right)
}

// 非递归
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) != 0 {
		for root != nil { //直到访问到根节点和所有左节点都访问完
			// 访问当前节点（根），放到结果里
			result = append(result, root.Val)
			// 将该节点（根）入栈
			stack = append(stack, root)
			// 访问左子树
			root = root.Left
		}
		// 弹出栈顶元素
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		//访问右子树
		root = node.Right
	}
	return result
}
