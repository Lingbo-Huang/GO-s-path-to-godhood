package interface_convert

import "testing"

func BenchmarkConvert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result := Convert(255) // 0-255不分配内存
		_ = result
	}
}

/*
go test -bench . -benchmem .
*/
