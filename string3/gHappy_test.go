package string3

import "testing"

func TestGHappy(t *testing.T) {
	if !gHappy("xxggxx") {
		t.Fatal("expected true but actually not")
	}
	if gHappy("xxgxx") {
		t.Fatal("expected false but actually not")
	}
	if gHappy("xxggyygxx") {
		t.Fatal("expected false but actually not")
	}
}
