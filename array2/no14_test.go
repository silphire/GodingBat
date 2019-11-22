package array2

import "testing"

func TestNo14(t *testing.T) {
	if !no14([]int{1, 2, 3}) {
		t.Fatal("expected true but actually not")
	}
	if no14([]int{1, 2, 3, 4}) {
		t.Fatal("expected false but actually not")
	}
	if !no14([]int{2, 3, 4}) {
		t.Fatal("expected true but actually not")
	}
}
