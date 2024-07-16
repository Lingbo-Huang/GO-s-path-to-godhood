package swap_nodes_in_pairs

/*
给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	return helper(head)
}

func helper(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 保存下一阶段的头指针
	nextHead := head.Next.Next
	// 翻转当前阶段指针的俩节点
	next := head.Next
	next.Next = head
	head.Next = helper(nextHead)
	return next
}
