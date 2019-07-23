package string1

import "testing"

func TestFrontAgain(t *testing.T) {
	if !frontAgain("edited") {
		t.Fatal("expected true but actually not")
	}
}

func TestNotFrontAgain(t *testing.T) {
	if frontAgain("played") {
		t.Fatal("expected true but actually not")
	}
}