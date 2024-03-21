package out

import "fmt"

// 自定义结构体Out，里面包含输出的数据，用的是任意类型的chan
type Out struct {
	data chan interface{}
}

var out *Out

func NewOut() *Out {
	if out == nil {
		out = &Out{
			data: make(chan interface{}, 65535),
		}
	}
	return out
}

func Println(i interface{}) {
	out = NewOut()
	out.data <- i
}

func (o *Out) OutPut() {
	for {
		select {
		case i := <-o.data:
			fmt.Println(i)
		}
	}
}
