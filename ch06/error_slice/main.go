package main

import "fmt"

func mPrint(datas ...interface{}) {
	for _, value := range datas {
		fmt.Println(value)
	}
}

func mPrint2(data interface{}) {
	fmt.Println(data)
}

func main() {

	var data = []interface{}{
		"bob", 18, 1.80,
	}
	mPrint(data...)

	var data1 = []string{
		"hh", "jsj", "sdfg",
	}
	//mPrint(data1...)
	//Cannot use 'data1' (type []string) as the type []interface{}
	var datai []interface{}
	for _, value := range data1 {
		datai = append(datai, value)
	}
	mPrint(datai...)
}
