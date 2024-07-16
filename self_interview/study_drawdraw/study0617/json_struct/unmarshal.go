package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Id   int
	Name string
	Age  int
}

func main() {

	var students []Student
	myClass :=
		`[
      {
        "id": 1,
        "name": "张三",
        "age": 20
      },
      {
        "id": 2,
        "name": "李翠花",
        "age": 18
      },
      {
        "id": 3,
        "name": "王老五",
        "age": 25
      }
  ]`

	err := json.Unmarshal([]byte(myClass), &students)

	if err == nil {
		for _, student := range students {
			fmt.Printf("%+v\n", student)
		}
	} else {
		fmt.Println(err)
	}
}
