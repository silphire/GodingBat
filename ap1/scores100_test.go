package ap1

import "testing"

func TestScores100(t *testing.T) {
	if !scores100([]int{1, 100, 100}) {
		t.Fatal("expected true but actual is not")
	}
	if scores100([]int{1, 100, 99, 100}) {
		t.Fatal("expected false but actual is not")
	}
	if !scores100([]int{100, 1, 100, 100}) {
		t.Fatal("expected true but actual is not")
	}
}
