package main

import (
	"fmt"
	"reflect"
)

func main() {
	/*
		报错，因为结构体类型虽然可以比较，但要求结构体里的字段顺序相同才行
		st1 := struct {
				age int
				name string
			}{age: 11, name:"hlb"}

			st2 := struct {
				name string
				age int
		}{age:11, name:"hlb"}
		if st1 == st2 {
			fmt.Println("st1 == st2")
		}
	*/
	/*
		map类型不可比较
			st3 := struct {
			age int
			m map[string]string
			}{age:11, m: map[string]string{"hlb":"oba"}}

			st4 := struct {
				age int
				m map[string]string
			}{age:11, m: map[string]string{"hlb":"oba"}}

			if st3 == st4 {
				fmt.Println("st3 == st4")
			}
	*/
	st3 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"hlb": "oba"}}

	st4 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"hlb": "oba"}}

	// 可以利用反射进行比较，里面其实是对类型和值分别比较
	if reflect.DeepEqual(st3, st4) {
		fmt.Println("st3 == st4")
	} else {
		fmt.Println("st3 != st4")
	}
}
