package logic2

import "testing"

func TestMakeChocolate(t *testing.T) {
	if makeChocolate(4, 1, 9) != 4 {
		t.Fatal("expected 4 but actually not")
	}
	if makeChocolate(4, 1, 10) != -1 {
		t.Fatal("expected -1 but actually not")
	}
	if makeChocolate(4, 1, 7) != 2 {
		t.Fatal("expected 2 but actually not")
	}
}
