package recursion1

import "testing"

func TestBunnyEars2(t *testing.T) {
	if bunnyEars2(0) != 0 {
		t.Fatal("expected 0 but actual is not")
	}
	if bunnyEars2(1) != 2 {
		t.Fatal("expected 2 but actual is not")
	}
	if bunnyEars2(2) != 5 {
		t.Fatal("expected 5 but actual is not")
	}
}
