package logic2

import "testing"

func TestNoTeenSum(t *testing.T) {
	if noTeenSum(1, 2, 3) != 6 {
		t.Fatal("expected 6 but actually not")
	}
	if noTeenSum(2, 13, 1) != 3 {
		t.Fatal("expected 3 but actually not")
	}
	if noTeenSum(2, 1, 14) != 3 {
		t.Fatal("expected 3 but actually not")
	}
}