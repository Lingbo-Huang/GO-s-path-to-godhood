package linear_list

import "errors"

/*
循环单链表
*/

// 循环单链表节点
type CircleSingleListNode struct {
	data interface{}
	next *CircleSingleListNode
}

// 循环单链表
type CircleSingleList struct {
	tail *CircleSingleListNode
	size int
}

// 初始化循环单链表
func InitCircleSingleList() *CircleSingleList {
	return &CircleSingleList{tail: nil, size: 0}
}

// 初始化循环单链表节点
// 这里需要传入data值，是因为循环单链表没有头节点这个说法，每个节点都要存储数据
// 这也是为什么要用两个结构体，一个节点，一个链表
func InitCircleSingleListNode(data interface{}) *CircleSingleListNode {
	return &CircleSingleListNode{data: data, next: nil}
}

// 单循环链表在表尾添加数据
func (csl *CircleSingleList) Append(data interface{}) {
	node := InitCircleSingleListNode(data)
	if csl.size == 0 {
		node.next = node
	} else {
		// 考虑特殊情况，本身只有一个节点时，插入
		// 本身节点其实就是tail指向的，它的next指向自身，也就是csl.tail.next == csl.tail
		// 这时候插入，下面这两句也能满足
		node.next = csl.tail.next
		csl.tail.next = node
	}
	csl.tail = node
	csl.size++
}

// 单循环链表查询数据, 返回下标为num的节点
func (csl *CircleSingleList) Get(num int) *CircleSingleListNode {
	if num < 0 || num > csl.size-1 {
		return nil
	}
	curNode := csl.tail
	for i := 0; i < num; i++ {
		curNode = curNode.next
	}
	return curNode
}

// 单循环链表插入数据data，使得新插入的节点位于从链表的tail开始遍历的下标为num处
func (csl *CircleSingleList) Insert(num int, data interface{}) error {
	node := InitCircleSingleListNode(data)
	if csl.size == 0 || csl.size == num {
		csl.Append(node)
	} else {
		var curNode *CircleSingleListNode
		if num == 0 {
			curNode = csl.tail
		} else {
			curNode = csl.Get(num)
		}
		//在curnode后面插入node
		node.next = curNode.next
		curNode.next = node
		csl.size++
	}
	return nil
}

// 单循环链表查询全部数据
func (csl *CircleSingleList) GetAll() []interface{} {
	result := make([]interface{}, 0)
	curNode := csl.tail
	for i := 0; i < csl.size; i++ {
		result = append(result, curNode.data)
		curNode = curNode.next
	}
	return result
}

// 单循环链表按序号删除数据
func (csl *CircleSingleList) RemoveInt(num int) error {
	if csl.size == 0 {
		return errors.New("循环链表为空")
	}
	if num > csl.size-1 {
		return errors.New("下标越界")
	}
	if csl.size == 1 {
		csl.tail = nil
		csl.size = 0
		return nil
	} else {
		var curNode *CircleSingleListNode
		var node *CircleSingleListNode
		if num == 0 {
			curNode = csl.tail
		} else {
			curNode = csl.Get(num - 1)
		}
		//curNode为num前一个节点（当num为0时就是链表指针）
		//node就是序号为num的节点
		node = curNode.next
		curNode.next = node.next
		if num == csl.size-1 {
			csl.tail = curNode
		}
		node.next = nil
		node = nil
		csl.size--
		return nil
	}
}

// 单循环链表删除全部元素
func (csl *CircleSingleList) RemoveAll() bool {
	if csl.size == 0 {
		return false
	}
	for i := 0; i < csl.size; i++ {
		curNode := csl.tail
		csl.tail = curNode.next
		curNode.next = nil
	}
	//for里没有将curNode = nil 是因为gc自动回收
	csl.tail = nil
	csl.size = 0
	return true
}
