package recursion1

import "testing"

func TestFibonacci(t *testing.T) {
	if fibonacci(0) != 0 {
		t.Fatal("expected 0 but actuak is not")
	}
	if fibonacci(1) != 1 {
		t.Fatal("expected 1 but actuak is not")
	}
	if fibonacci(2) != 1 {
		t.Fatal("expected 1 but actuak is not")
	}
}
