package detect_cycle

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
返回链表开始入环的第一个节点，如果链表无环，返回nil
*/

// // 思路：快慢指针，快慢相遇之后，慢指针从头走，快指针从下一个位置走
// 快慢指针步调一致一起移动，相遇点即为入环点
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	fast := head.Next
	slow := head
	for fast != nil && fast.Next != nil {
		if fast == slow {
			slow = head
			fast = fast.Next
			for fast != slow {
				fast = fast.Next
				slow = slow.Next
			}
			return slow
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	return nil
}
