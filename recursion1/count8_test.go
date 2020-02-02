package recursion1

import "testing"

func TestCount8(t *testing.T) {
	if count8(8) != 1 {
		t.Fatal("expected 1 but actual is not")
	}

	if count8(818) != 2 {
		t.Fatal("expected 2 but actual is not")
	}

	if count8(8818) != 4 {
		t.Fatal("expected 4 but actual is not")
	}
}
