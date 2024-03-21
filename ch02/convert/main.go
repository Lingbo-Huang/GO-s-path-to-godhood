package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a uint = 12
	var b uint8 = uint8(a)
	fmt.Println(b)

	var f float32 = 3.14
	var c = int32(f)
	fmt.Println(c)

	var f64 = float64(a)
	fmt.Println(f64)

	type IT int
	var d IT = IT(a)
	fmt.Println(d)

	var istr = "12"
	myint, err := strconv.Atoi(istr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(myint)

	//if s, err := strconv.Atoi(istr); err == nil {
	//	fmt.Printf("%T, %v \n", s, s)
	//}

	var myi = 8
	fmt.Println(strconv.Itoa(myi))

	//字符串转换为别的类型
	float, err := strconv.ParseFloat("3.1415", 64)
	if err != nil {
		return
	}
	fmt.Println(float)

	parseInt, err := strconv.ParseInt("12", 8, 64)
	if err != nil {
		return
	}
	fmt.Println(parseInt)

	parseBool, err := strconv.ParseBool("1")
	if err != nil {
		fmt.Println("ParseBool error")
		return
	}
	fmt.Println(parseBool)

	boolStr := strconv.FormatBool(true)
	fmt.Println(boolStr)

	floatStr := strconv.FormatFloat(3.1415, 'E', -1, 64)
	fmt.Println(floatStr)

	v := int64(-42)
	s10 := strconv.FormatInt(v, 10)
	fmt.Printf("%T, %v\n", s10, s10)
}
