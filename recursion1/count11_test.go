package recursion1

import "testing"

func TestCount11(t *testing.T) {
	if count11("11abc11") != 2 {
		t.Fatal("expected 1 but actual is not")
	}
	if count11("abc11x11x11") != 3 {
		t.Fatal("expected 3 but actual is not")
	}
	if count11("111") != 1 {
		t.Fatal("expected 1 but actual is not")
	}
}
