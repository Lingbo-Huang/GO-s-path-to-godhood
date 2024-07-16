package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

func printNode(node *Node) {
	for node.Next != nil {
		node = node.Next
		fmt.Println(node.Val)
	}
}

// 1 2 3 4
// pre 1 node 2->3->4
// pre 2->1 node 3->4
// pre 3->2->1 node 4
// pre 4->3->2->1
func linkReverse(node *Node) *Node {
	var pre *Node = nil
	for node != nil {
		temp := node.Next
		node.Next = pre
		pre = node
		node = temp
	}
	return pre
}

func main() {
	var header *Node = new(Node)
	tail := header
	for i := 1; i < 5; i++ {
		tail.Next = &Node{i, nil}
		tail = tail.Next
	}
	printNode(header)
	header.Next = linkReverse(header.Next)
	fmt.Println("反转之后")

}
