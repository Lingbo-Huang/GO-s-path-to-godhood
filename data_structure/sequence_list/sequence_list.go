package sequence_list

import "fmt"

// 数据结构之线性表--顺序表
type SqListInterFace interface {
	// 基本操作
	NewSqList(capacity int) *SqList // 初始化
	InitList(capacity int)          // 初始化
	ListEmpty() bool                // 判空
	ListFull() bool                 // 判满
	ListLength() int                // 返回数据元素个数
	ClearList()                     // 清空
	DestroyList()                   // 销毁
	// 元素操作
	ListInsert(index int, elem interface{}) bool // 插入元素
	ListDelete(index int) bool                   // 删除元素
	GetElem(index int) (interface{}, bool)       // 获取元素
	SetElem(elem interface{}, index int) bool    // 更新元素
	LocateElem(elem interface{}) (int, bool)     // 返回第1个值与elem相同的元素的位置若这样的数据元素不存在,则返回值为0
	// 其他操作
	PriorElem(elem interface{}) (interface{}, bool) // 寻找元素的前驱（当前元素的前一个元素）
	NextElem(elem interface{}) (interface{}, bool)  // 寻找元素的后驱（当前元素的后一个元素）
	TraverseList()                                  // 遍历
	Pop() interface{}                               // 从末尾弹出一个元素
	Append(elem interface{}) bool                   // 从末尾插入一个元素
	Reverse()                                       // 反转
}

// SqList 顺序表的结构类型为SqList
// 使用golang语言的interface接口类型创建顺序表
type SqList struct {
	Len         int
	Capacity    int
	Data        []interface{}
	ExtendRatio int
}

const NumberTwo int = 2

// NewSeqList 初始化
func (l *SqList) NewSqList(capacity int) *SqList {
	return &SqList{
		Len:         0,
		Capacity:    capacity,
		Data:        make([]interface{}, capacity),
		ExtendRatio: NumberTwo,
	}
}

// InitList 初始化
func (l *SqList) InitList(capacity int) {
	l.Capacity = capacity
	l.Len = 0
	m := make([]interface{}, capacity)
	l.Data = m
	l.ExtendRatio = NumberTwo
}

// ListEmpty 判空
func (l *SqList) ListEmpty() bool {
	if l.Len == 0 {
		return true
	} else {
		return false
	}
}

// ListLength 获取长度
func (l *SqList) ListLength() int {
	return l.Len
}

// ListFul 判满
func (l *SqList) ListFull() bool {
	if l.Len == l.Capacity {
		return true
	} else {
		return false
	}
}

// GetElem 根据下标Get元素
func (l *SqList) GetElem(index int) (interface{}, bool) {
	if index < 0 || index > l.Len {
		return nil, false
	} else {
		return l.Data[index], true
	}
}

// SetElem 更新元素
func (l *SqList) SetElem(elem interface{}, index int) bool {
	if index < 0 || index > l.Len {
		return false
	} else {
		l.Data[index] = elem
		return true
	}
}

// LocateELem 根据传入的值，返回第一个匹配的元素下标
func (l *SqList) LocateElem(elem interface{}) (int, bool) {
	for i, _ := range l.Data {
		if elem == l.Data[i] {
			return i, true
		}
	}
	return -1, false
}

// PriorElem 寻找元素的前驱（当前元素的前一个元素）
func (l *SqList) PriorElem(elem interface{}) (interface{}, bool) {
	i, _ := l.LocateElem(elem)
	if i == -1 || i == 0 {
		return nil, false
	} else {
		pre := l.Data[i-1]
		return pre, true
	}
}

// NextElem 寻找元素的后驱（当前元素的后一个元素）
func (l *SqList) NextElem(elem interface{}) (interface{}, bool) {
	i, _ := l.LocateElem(elem)
	if i == l.Len || i == -1 {
		return nil, false
	} else {
		next := l.Data[i+1]
		return next, true
	}
}

// ListInsert 插入元素,index为插入的位置，elem为插入值
func (l *SqList) ListInsert(index int, elem interface{}) bool {
	if index < 0 || index > l.Capacity || l.ListFull() {
		return false
	} else {
		for i := l.Len - 1; i >= index; i-- {
			l.Data[i+1] = l.Data[i]
		}
		l.Data[index] = elem
		l.Len++
		return true
	}
}

// ListDelete 删除元素
func (l *SqList) ListDelete(index int) bool {
	if index < 0 || index > l.Capacity || l.ListEmpty() {
		return false
	} else {
		for i := index; i < l.Len-1; i++ {
			l.Data[i] = l.Data[i+1]
		}
		l.Len--
		return true
	}
}

// TraverseList 遍历
func (l *SqList) TraverseList() {
	for i := 0; i < l.Len; i++ {
		fmt.Println(l.Data[i])
	}
	//下面这种遍历不可取，因为会把整个Capacity大小的空间里的值都打印出来
	//for _, val := range l.Data {
	//	fmt.Println(val)
	//}
}

// ClearList 清空
func (l *SqList) ClearList() {
	l.Len = 0
	l.Data = nil
}

// DestroyList 销毁
func (l *SqList) DestroyList() {
	l.Data = nil
	l.Len = 0
	l.Capacity = 0
	l.ExtendRatio = 0
}

// Pop 从末尾弹出一个元素
func (l *SqList) Pop() interface{} {
	if l.ListEmpty() {
		panic("线性表为空，没有可弹出的元素")
	}
	result := l.Data[l.Len-1]
	l.Data = l.Data[:l.Len-1]
	l.Len--
	return result
}

// Append 从末尾插入一个元素
func (l *SqList) Append(elem interface{}) bool {
	if l.ListFull() {
		panic("线性表已满，无法插入元素")
	}
	l.Data = append(l.Data, elem)
	l.Len++
	return true
}

// ExtendCapacity 扩容
func (l *SqList) ExtendCapacity() {
	// 因为append函数自身实现有扩容机制，所以，不需要考虑l.Data后面连续空间不足的问题，不足的话会分配新的内存空间创建新的切片
	l.Data = append(l.Data, make([]interface{}, l.Capacity*(l.ExtendRatio-1))...)
	l.Capacity = len(l.Data)
}

// Reverse 反转
func (l *SqList) Reverse() {
	for i := 0; i < l.Len>>1; i++ {
		l.Data[i], l.Data[l.Len-i-1] = l.Data[l.Len-i-1], l.Data[i]
	}
}
