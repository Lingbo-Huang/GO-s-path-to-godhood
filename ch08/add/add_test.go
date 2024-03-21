package add

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestAdd(t *testing.T) {
	re := add(1, 3)
	if re != 4 {
		t.Errorf("expect %d, actual %d", 4, re)
	}
}

func TestAdd2(t *testing.T) {
	fmt.Println("跳不跳过都执行")
	if testing.Short() {
		t.Skip("short 模式下跳过")
	}
	fmt.Println("没跳过")
	re := add(1, 2)
	if re != 3 {
		t.Errorf("expect %d, actual: %d", 3, re)
	}
}

func TestAdd3(t *testing.T) {
	var dataset = []struct {
		a   int
		b   int
		out int
	}{
		{1, 1, 2},
		{0, 0, 0},
		{-1, 1, 0},
		{-1, -8, -9},
	}

	for _, value := range dataset {
		re := add(value.a, value.b)
		if re != value.out {
			t.Errorf("expect: %d, actual: %d", value.out, re)
		}
	}
}

func BenchmarkAdd(bb *testing.B) {
	var a, b, c int
	a = 123
	b = 344
	c = 467

	for i := 0; i < bb.N; i++ {
		if actual := add(a, b); actual != c {
			fmt.Printf("%d + %d, expect:%d, actual:%d", a, b, c, actual)
		}
	}
}

const numbers = 10000

func BenchmarkStringSprintf(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var str string
		for j := 0; j < numbers; j++ {
			str = fmt.Sprintf("%s%d", str, j)
		}
	}

	b.StopTimer()
}

func BenchmarkStringAdd(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var str string
		for j := 0; j < numbers; j++ {
			str = str + strconv.Itoa(j)
		}
	}

	b.StopTimer()
}

func BenchmarkStringBuilder(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var str string
		var builder strings.Builder
		builder.WriteString(str)
		for j := 0; j < numbers; j++ {
			builder.WriteString(strconv.Itoa(j))
		}
		_ = builder.String()
	}

	b.StopTimer()
}
