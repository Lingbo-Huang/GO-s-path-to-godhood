package interface_convert

func Convert(val int) []interface{} {
	var list = make([]interface{}, 100)
	for i := 0; i < 100; i++ {
		list[i] = val
	}
	return list
}

/*
基本类型转interface{}类型，接口底层有一个指向类型的指针和一个指向值的指针。
在对int类型转换成interface时，0-255的数在静态数组里存好了，如果是0-255之间就不再转换了，不需要再申请内存
*/
