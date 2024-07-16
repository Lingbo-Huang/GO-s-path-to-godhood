package main

import "fmt"

func main() {
	/*
		常量位移表达式：操作数是常量(用len作用常量得到的还是常量，但作用于:这种切片就不是了)
		如果非常量位移表达式左操作数是无类型常量，那么会先隐式转换为整个位移表达式被左操作数单独替换后的类型
	*/
	const s = "CaiXiaobo"
	var a byte = 1 << len(s) / 128    // 4 s长度为9，2^9 / 2^7 = 4
	var b byte = 1 << len(s[:]) / 128 // 0 1 和 128 都变成byte类型，8 位, 0~255, 所以越界变成0
	fmt.Println(a)
	fmt.Println(b)
}
