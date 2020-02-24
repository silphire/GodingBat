package recursion2

import "testing"

func TestGroupSum6(t *testing.T) {
	if !groupSum6(0, []int{5, 6, 2}, 8) {
		t.Fatal("expected true but actual is not")
	}
	if groupSum6(0, []int{5, 6, 2}, 9) {
		t.Fatal("expected false but actual is not")
	}
	if groupSum6(0, []int{5, 6, 2}, 7) {
		t.Fatal("expected false but actual is not")
	}
}
