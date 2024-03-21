package main

import "fmt"

func main() {
	var mymap = map[string]string{
		"go":   "grpc",
		"cpp":  "stl",
		"java": "spring",
	}
	fmt.Println(mymap)
	var cMap = make(map[string]string, 3)
	cMap["hh"] = "mysql"
	fmt.Println(cMap)

	for key, value := range mymap {
		fmt.Println(key, value)
	}

	if _, ok := mymap["hlb"]; !ok {
		fmt.Println("not in")
	} else {
		fmt.Println("in")
	}

	delete(mymap, "cpp")
	fmt.Println(mymap)

}
