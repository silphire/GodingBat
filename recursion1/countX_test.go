package recursion1

import "testing"

func TestCountX(t *testing.T) {
	if countX("xxhixx") != 4 {
		t.Fatal("expected 4 but actual is not")
	}
	if countX("xhixhix") != 3 {
		t.Fatal("expected 3 but actual is not")
	}
	if countX("hi") != 0 {
		t.Fatal("expected 0 but actual is not")
	}
}
