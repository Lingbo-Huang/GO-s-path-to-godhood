package max_path_sum_binary_tree

/*
给定一个非空二叉树，返回其最大路径和。
最大路径和：在二叉树中找到一条路径，使得该路径上节点值的和最大。
最大路径经过当前节点：
		这个路径可以是当前节点的值，
		或者是当前节点的值加上左子树中的最大路径和，
		或者是当前节点的值加上右子树中的最大路径和，
		或者是当前节点的值加上左右子树中的最大路径和。
最大路径不经过当前节点：
		即当前节点的左子树或右子树中的最大路径和。
*/

/*
分治法，分为三种情况：左子树最大路径和最大，右子树最大路径和最大，左右子树最大加根节点最大，
需要保存两个变量：一个保存子树最大路径和，一个保存左右加根节点和，然后比较这个两个变量选择最大值即可
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ResultType struct {
	SinglePath int // 保持单边最大值
	MaxPath    int // 保存最大值（单边或者两个单边+根的值）
}

func maxPathSum(root *TreeNode) int {
	result := helper(root)
	return result.MaxPath
}

func helper(root *TreeNode) ResultType {
	if root == nil {
		return ResultType{
			SinglePath: 0,
			MaxPath:    -(1 << 31),
		}
	}

	left := helper(root.Left)
	right := helper(root.Right)

	result := ResultType{}
	if left.SinglePath > right.SinglePath {
		result.SinglePath = max(left.SinglePath+root.Val, 0)
	} else {
		result.SinglePath = max(right.SinglePath+root.Val, 0)
	}
	maxPath := max(left.MaxPath, right.MaxPath)
	result.MaxPath = max(maxPath, left.SinglePath+right.SinglePath+root.Val)
	return result
}
