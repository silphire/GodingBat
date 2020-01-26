package ap1

import "testing"

func TestCommonTwo(t *testing.T) {
	if commonTwo([]string{"a", "c", "x"}, []string{"b", "c", "d", "x"}) != 2 {
		t.Fatal("expected 2 but actual is not")
	}

	if commonTwo([]string{"a", "c", "x"}, []string{"a", "b", "c", "x", "z"}) != 3 {
		t.Fatal("expected 3 but actual is not")
	}

	if commonTwo([]string{"a", "b", "c"}, []string{"a", "b", "c"}) != 3 {
		t.Fatal("expected 3 but actual is not")
	}
}
