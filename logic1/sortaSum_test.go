package logic1

import "testing"

func TestSortaSum(t *testing.T) {
	if sortaSum(3, 4) != 7 {
		t.Fatal("expected 7 but actually not")
	}
	if sortaSum(9, 8) != 20 {
		t.Fatal("expected 20 but actually not")
	}
	if sortaSum(9, 13) != 22 {
		t.Fatal("expected 22 but actually not")
	}
}