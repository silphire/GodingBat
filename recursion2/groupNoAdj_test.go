package recursion2

import "testing"

func TestGroupNoAdj(t *testing.T) {
	if !groupNoAdj(0, []int{2, 5, 10, 4}, 12) {
		t.Fatal("expected true but actual is not")
	}
	if groupNoAdj(0, []int{2, 5, 10, 4}, 14) {
		t.Fatal("expected false but actual is not")
	}
	if groupNoAdj(0, []int{2, 5, 10, 4}, 7) {
		t.Fatal("expected false but actual is not")
	}
}
