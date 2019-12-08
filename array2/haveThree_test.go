package array2

import "testing"

func TestHaveThree(t *testing.T) {
	if !haveThree([]int{3, 1, 3, 1, 3}) {
		t.Fatal("expected true but actual is not")
	}

	if haveThree([]int{3, 1, 3, 3}) {
		t.Fatal("expected false but actual is not")
	}

	if haveThree([]int{3, 4, 3, 3, 4}) {
		t.Fatal("expected false but actual is not")
	}
}
