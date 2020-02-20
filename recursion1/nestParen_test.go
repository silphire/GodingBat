package recursion1

import "testing"

func TestNestParen(t *testing.T) {
	if !nestParen("(())") {
		t.Fatal("expected true but actual is not")
	}
	if !nestParen("((()))") {
		t.Fatal("expected true but actual is not")
	}
	if nestParen("(((x))") {
		t.Fatal("expected false but actual is not")
	}
}
