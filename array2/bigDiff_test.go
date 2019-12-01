package array2

import "testing"

func TestBigDiff(t *testing.T) {
	if bigDiff([]int{10, 3, 5, 6}) != 7 {
		t.Fatal("expected 7 but acutal is not")
	}
	if bigDiff([]int{7, 2, 10, 9}) != 8 {
		t.Fatal("expected 8 but acutal is not")
	}
	if bigDiff([]int{2, 10, 7, 2}) != 8 {
		t.Fatal("expected 8 but acutal is not")
	}
}
