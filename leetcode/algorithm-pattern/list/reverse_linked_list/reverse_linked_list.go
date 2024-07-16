package reverse_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
反转单链表
*/

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		// 用temp指向head的next
		temp := head.Next
		// 处理head指向的节点，把它用头插法放到prev里
		head.Next = prev
		prev = head
		// head往后走，也就是指向temp
		head = temp
	}
	return prev
}

/*
翻转从m到n的链表
*/

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	// 输入: 1->2->3->4->5->NULL, m = 2, n = 4
	if head == nil {
		return head
	}
	// 头部可能变化，使用dummy node
	dummy := &ListNode{Val: 0}
	dummy.Next = head
	head = dummy
	// 最开始：0->1->2->3->4->5->nil
	// 保存第m个元素的前面一个链表元素
	var pre *ListNode
	var i = 0
	for i < m {
		pre = head
		head = head.Next
		i++
	}
	// 遍历之后： 1(pre)->2(head)->3->4->5->NULL
	var j = i
	var next *ListNode
	var mid = head
	for head != nil && j <= n {
		// 第一次循环： 1(pre) 2(mid,next)->nil 3(head)->4->5->nil
		temp := head.Next
		head.Next = next
		next = head // 432
		head = temp // 5
		j++
	}
	// 循环需要执行三次
	// 循环结束：1(pre) 4(next)->3->2(mid)->nil 5(head)->nil
	pre.Next = next
	mid.Next = head
	return dummy.Next
}
