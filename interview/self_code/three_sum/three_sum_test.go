package three_sum

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestThreeSum(t *testing.T) {
	tests := []struct {
		input    []int
		expected [][]int
	}{
		{[]int{}, [][]int{}},
		{[]int{-1, 0, 1, 2, -1, 4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{[]int{0}, [][]int{}},
		{[]int{0, 0, 0, 0}, [][]int{{0, 0, 0}}},
	}

	for _, test := range tests {
		result := threeSum(test.input)
		if !equal(result, test.expected) {
			t.Errorf("threeSum(%v) = %v; expected %v", test.input, result, test.expected)
		}
	}
}

func equal(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		sort.Ints(a[i])
	}
	for i := range b {
		sort.Ints(b[i])
	}
	aMap := make(map[string]bool)
	bMap := make(map[string]bool)
	for _, triple := range a {
		aMap[tripleKey(triple)] = true
	}
	for _, triple := range b {
		bMap[tripleKey(triple)] = true
	}
	return reflect.DeepEqual(aMap, bMap)
}

func tripleKey(triple []int) string {
	return fmt.Sprintf("%d, %d, %d", triple[0], triple[1], triple[2])
}
