package binary_search

import "testing"

func TestSearch(t *testing.T) {
	nums := []int{
		1, 3, 4, 5, 5, 5, 5, 6, 7,
	}
	target := 5
	re := search(nums, target)
	t.Log(re)
}
