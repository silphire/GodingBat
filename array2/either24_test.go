package array2

import "testing"

func TestEither24(t *testing.T) {
	if !either24([]int{1, 2, 2}) {
		t.Fatal("expected true but actual is not")
	}
	if !either24([]int{4, 4, 1}) {
		t.Fatal("expected true but actual is not")
	}
	if either24([]int{4, 4, 1, 2, 2}) {
		t.Fatal("expected false but actual is not")
	}
}
