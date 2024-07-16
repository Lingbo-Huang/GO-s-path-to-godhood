package copy_random_list

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

/*
给定一个链表，每个节点包含一个额外增加的随机指针，
该指针可以指向链表中的任何节点或空节点。
要求返回这个链表的 深拷贝。
*/

/*
因为是深拷贝，所以直接复制节点是不行的，直接复制节点会使得指针指向原来的节点，那就是浅拷贝了
所以必须让Random指针指向新的复制后的节点
可以在原来的链表中每一个节点后复制一个节点，然后处理Random
元素复制完成，要分离两个链表
*/
func copyRandomList(head *Node) *Node {
	if head == nil {
		return head
	}

	// 把节点复制，放到复制的这个节点后面\
	// 1->2->3  ==>  1->1'->2->2'->3->3'
	cur := head
	for cur != nil {
		clone := &Node{Val: cur.Val, Next: cur.Next}
		temp := cur.Next
		cur.Next = clone
		cur = temp
	}
	// 处理random 指针
	cur = head
	for cur != nil {
		if cur.Random != nil {
			// 让复制的节点的Random指向原本Random指向的元素的复制体
			cur.Next.Random = cur.Random.Next
		}
		cur = cur.Next.Next
	}
	// 分离两个链表
	cur = head
	cloneHead := head.Next
	for cur != nil && cur.Next != nil {
		// cur本质上一次走一步，交叉构建两个链表
		temp := cur.Next
		// 串起来链表
		cur.Next = cur.Next.Next
		cur = temp
	}
	// 原始链表头：head 1->2->3
	// 克隆的链表头：cloneHead 1'->2'->3'
	return cloneHead
}
