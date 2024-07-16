package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Node struct {
	ID       string `json:"id"`
	Children []Node `json:"children"`
}

func printId(node Node) {
	fmt.Printf("id:%s\n", node.ID)
	for _, child := range node.Children {
		printId(child)
	}
}

func main() {
	data, err := os.ReadFile("./self_interview/test0603/data.json") // .表示项目路径
	if err != nil {
		fmt.Println("err:", err)
	}

	var root Node
	err = json.Unmarshal(data, &root)
	if err != nil {
		fmt.Println("err:", err)
	}
	//fmt.Println(string(data))
	//fmt.Println(root)
	//fmt.Printf("%+v", root)
	printId(root)
}
