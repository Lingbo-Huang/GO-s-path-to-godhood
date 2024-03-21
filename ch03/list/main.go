package main

import (
	"container/list"
	"fmt"
)

func main() {
	//var mylist list.List
	var mylist = list.New()

	mylist.PushBack("go")
	mylist.PushBack("grpc")
	mylist.PushBack("mysql")
	fmt.Println(mylist)

	for i := mylist.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}

	i := mylist.Front()
	for ; i != nil; i = i.Next() {
		if i.Value.(string) == "grpc" {
			break
		}
	}
	mylist.InsertBefore("gin", i)
	
	for i := mylist.Back(); i != nil; i = i.Prev() {
		fmt.Println(i.Value)
	}
}
