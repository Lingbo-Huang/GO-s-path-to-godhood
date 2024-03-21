package main

import "fmt"

type Node struct {
	data int
	next *Node
}

func Shownode(p *Node) { //遍历
	for p != nil {
		fmt.Println(*p)
		p = p.next //移动指针
	}
}

func main() {
	var head = new(Node)
	head.data = 0
	var tail *Node
	tail = head //tail用于记录头结点的地址，刚开始tail的的指针指向头结点
	for i := 1; i < 10; i++ {
		var node = Node{data: i}
		node.next = tail //将新插入的node的next指向头结点
		tail = &node     //重新赋值头结点
		//尾插法
		//(*tail).next = &node
		//tail = &node
	}

	Shownode(tail) //遍历结果
}
