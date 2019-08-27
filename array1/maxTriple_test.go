package array1

import "testing"

func TestMaxTriple(t *testing.T) {
	if maxTriple([]int{1, 2, 3}) != 3 {
		t.Fatal("expected 3 but actually not")
	}
	if maxTriple([]int{1, 5, 3}) != 5 {
		t.Fatal("expected 5 but actually not")
	}
	if maxTriple([]int{4, 2, 3}) != 4 {
		t.Fatal("expected 4 but actually not")
	}
}