package partition_list

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
给定一个链表和一个特定值 x，对链表进行分隔，
使得所有小于 x 的节点都在大于或等于 x 的节点之前。
*/

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return head
	}
	headDummy := &ListNode{Val: 0}
	tailDummy := &ListNode{Val: 0}
	tail := tailDummy
	headDummy.Next = head
	head = headDummy
	for head.Next != nil {
		if head.Next.Val < x {
			head = head.Next
		} else {
			// 将>x的节点放到另一个链表里
			t := head.Next
			head.Next = head.Next.Next
			tail.Next = t
			tail = tail.Next
		}
	}
	// 拼接两个链表
	tail.Next = nil
	head.Next = tailDummy.Next
	return headDummy.Next
}
