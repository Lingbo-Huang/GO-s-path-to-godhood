package generate_bst

/*
给定一个整数 n，生成所有由 1 ... n 为节点所组成的二叉搜索树。
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return generate(1, n)
}

func generate(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}
	ans := make([]*TreeNode, 0)
	for i := start; i <= end; i++ {
		lefts := generate(start, i-1)
		rights := generate(i+1, end)
		for j := 0; j < len(lefts); j++ {
			for k := 0; k < len(rights); k++ {
				root := &TreeNode{Val: i}
				root.Left = lefts[j]
				root.Right = rights[k]
				ans = append(ans, root)
			}
		}
	}
	return ans
}
