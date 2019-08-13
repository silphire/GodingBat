package array1

import "testing"

func TestNo23(t *testing.T) {
	if !no23([]int{4, 5}) {
		t.Fatal("expected true but not")
	}
	if no23([]int{4, 2}) {
		t.Fatal("expected false but not")
	}
	if no23([]int{3, 5}) {
		t.Fatal("expected false but not")
	}
}