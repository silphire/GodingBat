package recursion1

import "testing"

func TestTriangle(t *testing.T) {
	if triangle(0) != 0 {
		t.Fatal("expected 0 but actual is not")
	}

	if triangle(1) != 1 {
		t.Fatal("expected 1 but actual is not")
	}

	if triangle(2) != 3 {
		t.Fatal("expected 3 but actual is not")
	}
}
