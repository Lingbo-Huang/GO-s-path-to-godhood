package reorder_list

/*
给定一个单链表 L 的头节点 head ，单链表 L 表示为：

L0 → L1 → … → Ln - 1 → Ln
请将其重新排列后变为：

L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	// 找到中间位置
	mid := findMiddle(head)
	// 将后半部分翻转
	tail := reverseList(mid.Next)
	// 断开联系
	mid.Next = nil
	// 轮流插入两部分的元素，实现合并
	head = mergeTwoLists(head, tail)
}

func findMiddle(head *ListNode) *ListNode {
	fast := head.Next
	slow := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		temp := head.Next
		head.Next = prev
		prev = head
		head = temp
	}
	return prev
}

// 合并的链表中：轮流放l1和l2的元素，所以用一个flag不停的变true,false
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{Val: 0}
	head := dummy
	toggle := true
	for l1 != nil && l2 != nil {
		if toggle {
			head.Next = l1
			l1 = l1.Next
		} else {
			head.Next = l2
			l2 = l2.Next
		}
		toggle = !toggle
		head = head.Next
	}
	for l1 != nil {
		head.Next = l1
		head = head.Next
		l1 = l1.Next
	}
	for l2 != nil {
		head.Next = l2
		head = head.Next
		l2 = l2.Next
	}
	return dummy.Next
}
