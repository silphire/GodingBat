package string3

import "testing"

func TestCountTriple(t *testing.T) {
	if countTriple("abcXXXabc") != 1 {
		t.Fatal("expected 1 but actual is not")
	}
	if countTriple("xxxabyyyycd") != 3 {
		t.Fatal("expected 3 but actual is not")
	}
	if countTriple("a") != 0 {
		t.Fatal("expected 0 but actual is not")
	}
}
