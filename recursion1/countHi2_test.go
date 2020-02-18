package recursion1

import "testing"

func TestCountHi2(t *testing.T) {
	if countHi2("ahixhi") != 1 {
		t.Fatal("expected 1 but actual is not")
	}
	if countHi2("ahibhi") != 2 {
		t.Fatal("expected 2 but actual is not")
	}
	if countHi2("xhixhi") != 0 {
		t.Fatal("expected 0 but actual is not")
	}
}
