package linear_list

/*
单链表
*/
// 单链表接口
type SingleLister interface {
	GetFirst() *SingleList
	GetLast() *SingleList
	Length() (length int)
	Add(data interface{}) bool
	GetElem(index int) (interface{}, error)
	Delete(index int) (interface{}, bool)
}

// 单链表结点
type SingleList struct {
	Data interface{}
	Next *SingleList
}

// 初始化
func NewSingleList() *SingleList {
	return &SingleList{
		Data: nil, //数据根据自己数据类型设置零值
		Next: nil,
	}
}

// 返回第一个结点
func (this *SingleList) GetFirst() *SingleList {
	//从这句话可以看出我们使用的是带有头结点的单链表
	//链表为空返回nil
	if this.Next == nil {
		return nil
	}
	return this.Next
}

// 返回最后一个结点
func (this *SingleList) GetLast() *SingleList {
	// 空链表返回nil
	if this.Next == nil {
		return nil
	}
	point := this
	for point.Next != nil {
		point = point.Next
	}
	return point
}

// 获取单链表的长度
func (this *SingleList) Length() (length int) {
	length = 0
	point := this

	for point.Next != nil {
		length++
		point = point.Next
	}
	return
}

// 往单链表的末尾加一个元素
func (this *SingleList) Add(data interface{}) bool {
	point := this
	for point.Next != nil {
		point = point.Next
	}
	//这里注意用引用
	tmpSingle := &SingleList{Data: data}
	point.Next = tmpSingle
	return true
}

// 获取所有结点的值
func (this *SingleList) GetAll() []interface{} {
	result := make([]interface{}, 0)
	point := this
	for point.Next != nil {
		result = append(result, point.Data)
		point = point.Next
	}
	//最后一个节点不要忘记
	result = append(result, point.Data)
	return result
}

// 获取索引为index的结点
func (this *SingleList) GetElem(index int) *SingleList {
	if index < 0 || index > this.Length() {
		panic("Something is wrong with index's range, please check it")
		return nil
	}
	point := this
	for i := 0; i < index; i++ {
		point = point.Next
	}
	return point
}

// 删除索引为index的结点, 返回删除的节点里的值和操作成功与否的bool值
func (this *SingleList) Delete(index int) (interface{}, bool) {
	if index < 0 || index > this.Length() {
		panic("Something is wrong with index's range, please check it")
		return nil, false
	}
	point := this
	for i := 0; i < index-1; i++ {
		point = point.Next
	}
	deletePoint := point.Next
	point.Next = point.Next.Next
	return deletePoint, true
}
