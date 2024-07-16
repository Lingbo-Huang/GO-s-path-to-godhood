package validate_bst

/*
验证二叉搜索树
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return dfs(root).valid
}

type ResultType struct {
	max   int
	min   int
	valid bool
}

func dfs(root *TreeNode) (result ResultType) {
	if root == nil {
		result.max = -1 << 63
		result.min = 1<<63 - 1
		result.valid = true
		return
	}
	left := dfs(root.Left)
	right := dfs(root.Right)
	if root.Val > left.max && root.Val < right.min && left.valid && right.valid {
		result.valid = true
	}
	// 更新当前节点的最大最小值
	result.max = Max(Max(left.max, right.max), root.Val)
	result.min = Min(Min(left.min, right.min), root.Val)
	return
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
