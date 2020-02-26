package recursion2

import "testing"

func TestGroupSumClump(t *testing.T) {
	if !groupSumClump(0, []int{2, 4, 8}, 10) {
		t.Fatal("expected true but actual is not")
	}
	if !groupSumClump(0, []int{1, 2, 4, 8, 1}, 14) {
		t.Fatal("expected true but actual is not")
	}
	if groupSumClump(0, []int{2, 4, 4, 8}, 14) {
		t.Fatal("expected false but actual is not")
	}
}
