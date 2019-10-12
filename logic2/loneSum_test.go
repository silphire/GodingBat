package logic2

import "testing"

func TestLoneSum(t *testing.T) {
	if loneSum(1, 2, 3) != 6 {
		t.Fatal("expected 6 but actually not")
	}
	if loneSum(3, 2, 3) != 2 {
		t.Fatal("expected 2 but actually not")
	}
	if loneSum(3, 3, 3) != 0 {
		t.Fatal("expected 0 but actually not")
	}
}
