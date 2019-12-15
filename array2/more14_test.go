package array2

import "testing"

func TestMore14(t *testing.T) {
	if !more14([]int{1, 4, 1}) {
		t.Fatal("expected true but actual is not")
	}
	if more14([]int{1, 4, 1, 4}) {
		t.Fatal("expected false but actual is not")
	}
	if !more14([]int{1, 1}) {
		t.Fatal("expected true but actual is not")
	}
}
