package level_order_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
返回二叉树按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）
*/
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	result := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		temp := make([]int, 0)
		length := len(queue)
		for length != 0 {
			value := queue[0]
			if value.Left != nil {
				queue = append(queue, value.Left)
			}
			if value.Right != nil {
				queue = append(queue, value.Right)
			}
			queue = queue[1:]
			length--
			temp = append(temp, value.Val)
		}
		result = append(result, temp)
	}
	return result
}
