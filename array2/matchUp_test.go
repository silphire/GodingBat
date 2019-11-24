package array2

import "testing"

func TestMatchUp(t *testing.T) {
	if matchUp([]int{1, 2, 3}, []int{2, 3, 10}) != 2 {
		t.Fatal("expected 2 but actual is not")
	}
	if matchUp([]int{1, 2, 3}, []int{2, 3, 5}) != 3 {
		t.Fatal("expected 3 but actual is not")
	}
	if matchUp([]int{1, 2, 3}, []int{2, 3, 3}) != 2 {
		t.Fatal("expected 2 but actual is not")
	}
}
