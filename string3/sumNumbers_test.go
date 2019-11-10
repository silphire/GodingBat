package string3

import "testing"

func TestSumNumbers(t *testing.T) {
	if sumNumbers("abc123xyz") != 123 {
		t.Fatal("expected 123 but actually not")
	}

	if sumNumbers("aa11b33") != 44 {
		t.Fatal("expected 44 but actually not")
	}

	if sumNumbers("7 11") != 18 {
		t.Fatal("expected 18 but actually not")
	}
}
