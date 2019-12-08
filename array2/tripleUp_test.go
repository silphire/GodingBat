package array2

import "testing"

func TestTripleUp(t *testing.T) {
	if !tripleUp([]int{1, 4, 5, 6, 2}) {
		t.Fatal("expected true but actual is not")
	}
	if !tripleUp([]int{1, 2, 3}) {
		t.Fatal("expected true but actual is not")
	}
	if tripleUp([]int{1, 2, 4}) {
		t.Fatal("expected false but actual is not")
	}
}
