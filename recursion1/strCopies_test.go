package recursion1

import "testing"

func TestStrCopies(t *testing.T) {
	if !strCopies("catcowcat", "cat", 2) {
		t.Fatal("expected true but actual is not")
	}
	if strCopies("catcowcat", "cow", 2) {
		t.Fatal("expected false but actual is not")
	}
	if !strCopies("catcowcat", "cow", 1) {
		t.Fatal("expected true but actual is not")
	}
}
