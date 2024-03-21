package main

import (
	"fmt"
)

func main() {
	//var i int
	//for i < 3 {
	//	time.Sleep(2 * time.Second)
	//	fmt.Println(i)
	//	i++
	//}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")

	var sum int
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d ", i, j, i*j)
		}
		fmt.Println()
	}

	name := "hhha, hlb, imooc"
	for key, value := range name {
		fmt.Printf("%d, %c\r\n", key, value)
	}

	map1 := make(map[int]string)
	map1[1] = "hlb"
	map1[2] = "love"
	map1[3] = "go"
	for key, value := range map1 {
		fmt.Printf("key is %d, value is %s\n", key, value)
	}
	for _, v := range map1 {
		fmt.Println(v)
	}

	var i, j int
	for i = 2; i < 100; i++ {
		for j = 2; j <= (i / j); j++ {
			if (i % j) == 0 {
				break
			}
		}
		if j > (i / j) {
			fmt.Printf("%d 是素数\n", i)
		}
	}
}
