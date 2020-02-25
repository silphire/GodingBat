package recursion2

import "testing"

func TestGroupSum5(t *testing.T) {
	if !groupSum5(0, []int{2, 5, 10, 4}, 19) {
		t.Fatal("expected false but actual is not")
	}
	if !groupSum5(0, []int{2, 5, 10, 4}, 17) {
		t.Fatal("expected false but actual is not")
	}
	if groupSum5(0, []int{2, 5, 10, 4}, 12) {
		t.Fatal("expected true but actual is not")
	}
}
