package main

import "fmt"

// Node 表示二叉树的节点
type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

// BinaryTree 表示二叉树的结构
type BinaryTree struct {
	Root *Node
}

// Insert 插入节点到二叉树
func (bt *BinaryTree) Insert(key int) {
	if bt.Root == nil {
		bt.Root = &Node{Key: key}
	} else {
		bt.insertRecursive(bt.Root, key)
	}
}

// insertRecursive 递归插入节点
func (bt *BinaryTree) insertRecursive(root *Node, key int) {
	if key < root.Key {
		if root.Left == nil {
			root.Left = &Node{Key: key}
		} else {
			bt.insertRecursive(root.Left, key)
		}
	} else {
		if root.Right == nil {
			root.Right = &Node{Key: key}
		} else {
			bt.insertRecursive(root.Right, key)
		}
	}
}

// InOrderTraversal 中序遍历二叉树
func (bt *BinaryTree) InOrderTraversal() {
	bt.inOrderTraversalRecursive(bt.Root)
	fmt.Println()
}

// inOrderTraversalRecursive 递归中序遍历
func (bt *BinaryTree) inOrderTraversalRecursive(root *Node) {
	if root != nil {
		bt.inOrderTraversalRecursive(root.Left)
		fmt.Printf("%d ", root.Key)
		bt.inOrderTraversalRecursive(root.Right)
	}
}

// PreOrderTraversal 前序遍历二叉树
func (bt *BinaryTree) PreOrderTraversal() {
	bt.preOrderTraversalRecursive(bt.Root)
	fmt.Println()
}

// preOrderTraversalRecursive 递归前序遍历
func (bt *BinaryTree) preOrderTraversalRecursive(root *Node) {
	if root != nil {
		fmt.Printf("%d ", root.Key)
		bt.preOrderTraversalRecursive(root.Left)
		bt.preOrderTraversalRecursive(root.Right)
	}
}

// PostOrderTraversal 后序遍历二叉树
func (bt *BinaryTree) PostOrderTraversal() {
	bt.postOrderTraversalRecursive(bt.Root)
	fmt.Println()
}

// postOrderTraversalRecursive 递归后序遍历
func (bt *BinaryTree) postOrderTraversalRecursive(root *Node) {
	if root != nil {
		bt.postOrderTraversalRecursive(root.Left)
		bt.postOrderTraversalRecursive(root.Right)
		fmt.Printf("%d ", root.Key)
	}
}

func main() {
	// 创建二叉树
	binaryTree := &BinaryTree{}

	// 插入节点
	binaryTree.Insert(50)
	binaryTree.Insert(30)
	binaryTree.Insert(70)
	binaryTree.Insert(20)
	binaryTree.Insert(40)
	binaryTree.Insert(60)
	binaryTree.Insert(80)

	// 中序遍历
	fmt.Print("InOrder Traversal: ")
	binaryTree.InOrderTraversal()

	// 前序遍历
	fmt.Print("PreOrder Traversal: ")
	binaryTree.PreOrderTraversal()

	// 后序遍历
	fmt.Print("PostOrder Traversal: ")
	binaryTree.PostOrderTraversal()
}
