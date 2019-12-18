package array2

import "testing"

func TestTwoTwo(t *testing.T) {
	if !twoTwo([]int{4, 2, 2, 3}) {
		t.Fatal("expected true but actual is not")
	}
	if !twoTwo([]int{2, 2, 4}) {
		t.Fatal("expected true but actual is not")
	}
	if twoTwo([]int{2, 2, 4, 2}) {
		t.Fatal("expected false but actual is not")
	}
}
