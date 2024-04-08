package traverse

import "fmt"

type ListNode struct {
	val  int
	next *ListNode
}

func traverseList(head *ListNode) {
	for p := head; p != nil; p = p.next {
		fmt.Println(p.val)
	}
}
