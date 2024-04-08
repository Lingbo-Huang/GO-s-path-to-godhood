package sequence_list

import (
	"fmt"
	"testing"
)

func TestSqList(t *testing.T) {
	var sl SqList
	//初始化
	sl.InitList(4)
	//sl = NewSqList(4)
	//判空
	fmt.Println(sl.ListEmpty()) // true
	//判满
	fmt.Println(sl.ListFull()) // false
	// 定义一个Struct类型
	type s struct {
		name string
		age  int
	}
	student1 := s{name: "hlb", age: 20}
	student2 := s{name: "ljx", age: 18}
	// 插入元素
	sl.ListInsert(0, student1)
	sl.ListInsert(1, student2)
	// 判空
	fmt.Println(sl.ListEmpty()) // false
	// 插入元素
	sl.ListInsert(2, 1000)
	sl.ListInsert(3, "GOGO")

	// 遍历
	sl.TraverseList() // {hlb 20} {ljx 18} 1000 GoGO
	// 获取长度
	fmt.Println(sl.ListLength()) // 4
	// 判满
	fmt.Println(sl.ListFull()) // true
	// 插入元素
	fmt.Println(sl.ListInsert(4, "jjj")) // false

	// 扩容
	sl.ExtendCapacity()
	// 获取长度
	fmt.Println("扩容后的容量：", sl.Capacity)  // 8
	fmt.Println(sl.ListInsert(4, "jjj")) // true
	// 遍历
	sl.TraverseList() // {hlb 20} {ljx 18} 1000 GoGO jjj

	// 删除元素,索引为2
	sl.ListDelete(2)
	// 遍历
	sl.TraverseList() // {hlb 20} {ljx 18} GoGO jjj

	// 根据下标Get元素
	el, _ := sl.GetElem(1)
	fmt.Println(el) // {ljx 18}

	// 更新元素
	fmt.Println("更新元素：", sl.SetElem("超哥哥", 2)) // true
	// 遍历
	sl.TraverseList() // {hlb 20} {ljx 18} 超哥哥 jjj

	// 根据传入的值，返回第一个匹配的元素下标
	b, b1 := sl.LocateElem(student2)
	fmt.Println(b, b1) // 1 true

	// 寻找元素的后驱
	n1, n2 := sl.NextElem(student2)
	fmt.Println(n1, n2) // 超哥哥 true

	// 寻找元素的前驱
	p1, p2 := sl.PriorElem("超哥哥")
	fmt.Println(p1, p2) // {ljx 18} true

	// 从末尾弹出一个元素
	p1 = sl.Pop()
	fmt.Println("从末尾弹出一个元素:", p1) // 从末尾弹出一个元素: jjj
	sl.TraverseList()             // 遍历 {hlb 20} {ljx 18} 超哥哥

	// 从末尾插入一个元素
	fmt.Println("从末尾插入一个元素", sl.Append("gopher")) // true
	sl.TraverseList()                             // 遍历 {hlb 20} {ljx 18} 超哥哥 gopher

	// 反转
	sl.Reverse()
	sl.TraverseList() // 遍历 gopher 超哥哥 {ljx 18} {hlb 20}

	// 清空
	sl.ClearList()
	fmt.Println(sl.ListEmpty()) // true

	// 遍历
	sl.TraverseList()
}
