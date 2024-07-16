package main

import "fmt"

type Node struct {
	value int
	Next  *Node
}

func main() {
	var n *Node = new(Node)
	n.value = 1
	n1 := new(Node)
	n1.value = 2
	n2 := new(Node)
	n2.value = 3
	n.Next = n1
	n1.Next = n2
	var preN *Node = nil
	for n != nil {
		preN, n, n.Next = n, n.Next, preN
	}
	for preN != nil {
		fmt.Println(preN.value)
		preN = preN.Next
	}

}
