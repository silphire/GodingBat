package recursion1

import "testing"

func TestPowerN(t *testing.T) {
	if powerN(3, 1) != 3 {
		t.Fatal("expected 3 but actual is not")
	}
	if powerN(3, 2) != 9 {
		t.Fatal("expected 9 but actual is not")
	}
	if powerN(3, 3) != 27 {
		t.Fatal("expected 27 but actual is not")
	}
}
