package logic1

import "testing"

func TestWithoutDoubles(t *testing.T) {
	if withoutDoubles(2, 3, true) != 5 {
		t.Fatal("expected 5 but actually not")
	}
	if withoutDoubles(3, 3, true) != 7 {
		t.Fatal("expected 7 but actually not")
	}
	if withoutDoubles(3, 3, false) != 6 {
		t.Fatal("expected 6 but actually not")
	}
}