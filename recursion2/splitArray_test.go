package recursion2

import "testing"

func TestSplitArray(t *testing.T) {
	if !splitArray([]int{2, 2}) {
		t.Fatal("expected true but actual is not")
	}
	if splitArray([]int{2, 3}) {
		t.Fatal("expected false but actual is not")
	}
	if !splitArray([]int{5, 2, 3}) {
		t.Fatal("expected true but actual is not")
	}
}
