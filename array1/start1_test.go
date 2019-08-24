package array1

import "testing"

func TestStart1(t *testing.T) {
	if start1([]int{1, 2, 3}, []int{1, 2}) != 2 {
		t.Fatal("expected 2 but actually not")
	}
	if start1([]int{1, 2, 3}, []int{}) != 1 {
		t.Fatal("expected 1 but actually not")
	}
}