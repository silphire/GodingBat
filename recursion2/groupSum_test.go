package recursion2

import "testing"

func TestGroupSum(t *testing.T) {
	if !groupSum(0, []int{2, 4, 8}, 10) {
		t.Fatal("expected true but actual is not")
	}
	if !groupSum(0, []int{2, 4, 8}, 14) {
		t.Fatal("expected true but actual is not")
	}
	if groupSum(0, []int{2, 4, 8}, 9) {
		t.Fatal("expected false but actual is not")
	}
}
