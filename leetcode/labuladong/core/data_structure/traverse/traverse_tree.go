package traverse

// 二叉树
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// n叉树
type TreeNodes struct {
	val      int
	children []*TreeNodes
}

func traverseTree(root *TreeNode) {
	if root != nil {
		traverseTree(root.Left)
		traverseTree(root.Right)
	}
}

func traverseTrees(root *TreeNodes) {
	if root != nil {
		for _, child := range root.children {
			traverseTrees(child)
		}
	}
}
