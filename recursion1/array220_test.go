package recursion1

import "testing"

func TestArray220(t *testing.T) {
	if !array220([]int{1, 2, 20}, 0) {
		t.Fatal("expected true but actual is not")
	}

	if !array220([]int{3, 30}, 0) {
		t.Fatal("expected true but actual is not")
	}

	if array220([]int{3}, 0) {
		t.Fatal("expected false but actual is not")
	}
}
