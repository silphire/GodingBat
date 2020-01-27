package recursion1

import "testing"

func TestBunnyEars(t *testing.T) {
	if bunnyEars(0) != 0 {
		t.Fatal("expected 0 but actual is not")
	}
	if bunnyEars(1) != 2 {
		t.Fatal("expected 2 but actual is not")
	}
	if bunnyEars(2) != 4 {
		t.Fatal("expected 4 but actual is not")
	}
}
