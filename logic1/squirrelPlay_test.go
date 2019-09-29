package logic1

import "testing"

func TestSquirrelPlay(t *testing.T) {
	if !squirrelPlay(70, false) {
		t.Fatal("expected true but actually not")
	}
	if squirrelPlay(95, false) {
		t.Fatal("expected false but actually not")
	}
	if !squirrelPlay(95, true) {
		t.Fatal("expected true but actually not")
	}
}