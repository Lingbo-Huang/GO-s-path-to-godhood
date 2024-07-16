package implement_queue_using_stacks

/* 题目
https://leetcode-cn.com/problems/implement-queue-using-stacks/
使用栈实现队列
*/

type MyQueue struct {
	stack []int
	back  []int
}

func Constructor() *MyQueue {
	return &MyQueue{
		stack: make([]int, 0),
		back:  make([]int, 0),
	}
}

// Push 入队
func (this *MyQueue) Push(x int) {
	// 把元素 x 压入栈 stack 的栈顶
	this.stack = append(this.stack, x)
}

// 取出 stack 的栈里所有元素，压到 back 栈里，使得数据顺序反过来
func (this *MyQueue) in2out() {
	for len(this.stack) != 0 {
		// 弹出栈 stack 的栈顶元素， 压入 back 栈顶
		val := this.stack[len(this.stack)-1]
		this.stack = this.stack[:len(this.stack)-1]
		this.back = append(this.back, val)
	}
}

// Pop 出队
func (this *MyQueue) Pop() int {
	if len(this.back) == 0 {
		this.in2out()
	}
	// 弹出栈 back 的栈顶元素，并返回
	val := this.back[len(this.back)-1]
	this.back = this.back[:len(this.back)-1]
	return val
}

// Peek 取队首元素
func (this *MyQueue) Peek() int {
	if len(this.back) == 0 {
		this.in2out()
	}
	// 取栈 back 的栈顶元素，并返回
	val := this.back[len(this.back)-1]
	return val
}

// Empty 判空
func (this *MyQueue) Empty() bool {
	return len(this.stack) == 0 && len(this.back) == 0
}
