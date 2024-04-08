package traverse

import "fmt"

func traverseArr(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}
