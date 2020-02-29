package recursion2

import "testing"

func TestSplit53(t *testing.T) {
	if !split53([]int{1, 1}) {
		t.Fatal("expected true but actual is not")
	}
	if split53([]int{1, 1, 1}) {
		t.Fatal("expected false but actual is not")
	}
	if !split53([]int{2, 4, 2}) {
		t.Fatal("expected true but actual is not")
	}
}
