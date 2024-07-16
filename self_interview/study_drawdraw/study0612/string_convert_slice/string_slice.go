package main

import (
	"fmt"
	"unsafe"
)

/*
强制类型转换一定会发生内存拷贝，因为不同类型在内存中存储方式不同
但是是可以通过指针的强转避免拷贝的
*/
func main() {
	// 字符串转切片
	s := "abcd"
	strData := unsafe.StringData(s)            // *byte
	byteSlice := unsafe.Slice(strData, len(s)) // []byte
	fmt.Printf("%#v\n", byteSlice)
	// 切片转字符串
	bs := []byte{'h', 'e', 'l', 'l', 'o'}
	str := unsafe.String(&bs[0], len(bs))
	fmt.Printf("%#v\n", bs)
	fmt.Printf("%#v\n", str)
}
