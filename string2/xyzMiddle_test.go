package string2

import "testing"

func TestXyzMiddle(t *testing.T) {
	if !xyzMiddle("AAxyzBB") {
		t.Fatal("expected true but actually not")
	}

	if !xyzMiddle("AxyzBB") {
		t.Fatal("expected true but actually not")
	}

	if xyzMiddle("AxyzBBB") {
		t.Fatal("expected false but actually not")
	}
}
