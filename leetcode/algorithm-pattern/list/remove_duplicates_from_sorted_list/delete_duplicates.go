package remove_duplicates_from_sorted_list

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。
*/
func delete_duplicates(head *ListNode) *ListNode {
	current := head
	for current != nil {
		for current.Next != nil && current.Val == current.Next.Val {
			current.Next = current.Next.Next
		}
		current = current.Next
	}
	return head
}

/*
给定一个排序链表，删除所有含有重复数字的节点，
只保留原始链表中 没有重复出现的数字。
*/
func delete_duplicates2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	// 头节点可能被删除，所以引入dummy node辅助删除
	dummy := &ListNode{Val: 0}
	dummy.Next = head
	head = dummy
	// 记录重复的值
	var rmVal int
	for head.Next != nil && head.Next.Next != nil {
		if head.Next.Val == head.Next.Next.Val {
			rmVal = head.Next.Val
			for head.Next != nil && head.Next.Val == rmVal {
				head.Next = head.Next.Next
			}
		} else {
			head = head.Next
		}
	}
	return dummy.Next
}
