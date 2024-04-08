package main

// 根据前序遍历和中序遍历的结果还原一棵二叉树
func build(preorder []int, preStart int, preEnd int,
	inorder []int, inStart int, inEnd int) *TreeNode {
	if preStart > preEnd {
		return nil
	}
	rootVal := preorder[preStart]
	index := 0
	for i := inStart; i < inEnd; i++ {
		if rootVal == inorder[i] {
			index = i
			break
		}
	}
	leftSize := index - inStart
	root := &TreeNode{Val: rootVal}

	root.Left = build(preorder, preStart+1, preStart+leftSize,
		inorder, inStart, index-1)
	root.Right = build(preorder, preStart+leftSize+1, preEnd,
		inorder, index+1, inEnd)
	return root
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	rootVal := preorder[0]
	index := 0
	for i := 0; i < len(inorder); i++ {
		if rootVal == inorder[i] {
			index = i
			break
		}
	}
	root := &TreeNode{Val: rootVal}
	leftPreorder := preorder[1 : index+1]
	leftInorder := inorder[:index]
	rightPreorder := preorder[index+1:]
	rightInorder := inorder[index+1:]

	root.Left = buildTree(leftPreorder, leftInorder)
	root.Right = buildTree(rightPreorder, rightInorder)
	return root
}
