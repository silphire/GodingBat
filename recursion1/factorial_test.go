package recursion1

import "testing"

func TestFactorial(t *testing.T) {
	if factorial(1) != 1 {
		t.Fatal("expected 1 but actual is not")
	}
	if factorial(2) != 2 {
		t.Fatal("expected 2 but actual is not")
	}
	if factorial(3) != 6 {
		t.Fatal("expected 6 but actual is not")
	}
}
