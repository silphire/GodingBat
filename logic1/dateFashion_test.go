package logic1

import "testing"

func TestDateFashion(t *testing.T) {
	if dateFashion(5, 10) != 2 {
		t.Fatal("expected 2 but actually not")
	}
	if dateFashion(1, 10) != 0 {
		t.Fatal("expected 0 but actually not")
	}
	if dateFashion(7, 3) != 1 {
		t.Fatal("expected 1 but actually not")
	}
}