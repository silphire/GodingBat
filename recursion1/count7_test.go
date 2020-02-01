package recursion1

import "testing"

func TestCount7(t *testing.T) {
	if count7(717) != 2 {
		t.Fatal("expected 2 but actual is not")
	}
	if count7(7) != 1 {
		t.Fatal("expected 1 but actual is not")
	}
	if count7(123) != 0 {
		t.Fatal("expected 0 but actual is not")
	}
}
