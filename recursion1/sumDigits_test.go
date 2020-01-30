package recursion1

import "testing"

func TestSumDigits(t *testing.T) {
	if sumDigits(126) != 9 {
		t.Fatal("expected 9 but actual is not")
	}

	if sumDigits(49) != 13 {
		t.Fatal("expected 13 but actual is not")
	}

	if sumDigits(12) != 3 {
		t.Fatal("expected 3 but actual is not")
	}
}
