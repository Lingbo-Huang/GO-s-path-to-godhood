package sort_list

/*
在O(nlogn)时间复杂度和常熟空间复杂度，对链表排序
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	return mergeSort(head)
}

// 归并排序
func mergeSort(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 找到中间节点
	middle := findMiddle(head)
	// tail指针指向后面的节点
	tail := middle.Next
	// 前半部分和后半部分分开
	middle.Next = nil
	// 前后分别归并排序
	left := mergeSort(head)
	right := mergeSort(tail)
	// 将前后合并
	result := merge(left, right)
	return result
}

func merge(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{Val: 0}
	head := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			head.Next = l1
			l1 = l1.Next
		} else {
			head.Next = l2
			l2 = l2.Next
		}
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

func findMiddle(head *ListNode) *ListNode {
	slow := head
	fast := head.Next
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}
