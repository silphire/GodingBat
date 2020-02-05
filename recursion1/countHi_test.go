package recursion1

import "testing"

func TestCountHi(t *testing.T) {
	if countHi("xxhixx") != 1 {
		t.Fatal("expected 1 but actual is not")
	}
	if countHi("xhixhix") != 2 {
		t.Fatal("expected 2 but actual is not")
	}
	if countHi("hi") != 1 {
		t.Fatal("expected 1 but actual is not")
	}
}
