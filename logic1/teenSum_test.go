package logic1

import "testing"

func TestTeenSum(t *testing.T) {
	if teenSum(3, 4) != 7 {
		t.Fatal("expected 7 but actually not")
	}
	if teenSum(10, 13) != 19 {
		t.Fatal("expected 19 but actually not")
	}
	if teenSum(13, 2) != 19 {
		t.Fatal("expected 19 but actually not")
	}
}
