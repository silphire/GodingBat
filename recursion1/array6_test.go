package recursion1

import "testing"

func TestArray6(t *testing.T) {
	if !array6([]int{1, 6, 4}, 0) {
		t.Fatal("expected true but actual is not")
	}
	if array6([]int{1, 4}, 0) {
		t.Fatal("expected false but actual is not")
	}
	if !array6([]int{6}, 0) {
		t.Fatal("expected false but actual is not")
	}
}
