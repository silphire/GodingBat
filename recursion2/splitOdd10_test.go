package recursion2

import "testing"

func TestSplitOdd10(t *testing.T) {
	if !splitOdd10([]int{5, 5, 5}) {
		t.Fatal("expected true but actual is not")
	}
	if splitOdd10([]int{5, 5, 6}) {
		t.Fatal("expected false but actual is not")
	}
	if !splitOdd10([]int{5, 5, 6, 1}) {
		t.Fatal("expected true but actual is not")
	}
}
