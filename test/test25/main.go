package main

/*
当 map 或 slice 作为函数参数传入时或者作为返回值传出时，
如果您存储了对它们的引用，则用户可以对其进行修改。
*/
import (
	"fmt"
	"sync"
)

type Driver struct {
	mu        sync.Mutex
	sliceData []string
	mapData   map[string]string
}

func (d *Driver) SetSliceData(data []string) {
	d.sliceData = make([]string, len(data))
	copy(d.sliceData, data)
}

func (d *Driver) GetSliceData() []string {
	d.mu.Lock()
	defer d.mu.Unlock()

	result := make([]string, len(d.sliceData))
	for k, v := range d.sliceData {
		result[k] = v
	}
	return result
}

func main() {
	testData := []string{"hlb", "love", "code"}
	d1 := new(Driver)
	d1.SetSliceData(testData)
	for _, v := range d1.sliceData {
		fmt.Printf("%v ", v)
	}
	fmt.Println()
	testData[0] = "ljx"
	for _, v := range d1.sliceData {
		fmt.Printf("%v ", v)
	}
	fmt.Println()

	testShot := d1.GetSliceData()
	for _, v := range testShot {
		fmt.Printf("%v ", v)
	}
	fmt.Println()

	testShot[0] = "ljx"
	for _, v := range d1.sliceData {
		fmt.Printf("%v ", v)
	}
	fmt.Println()
}
