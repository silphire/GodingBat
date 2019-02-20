package warmup1

import "testing"

func TestBothSmile(t *testing.T) {
	if !monkeyTrouble(true, true) {
		t.Fatal("expected we have trouble but actually not")
	}
}
func TestNotBothSmile(t *testing.T) {
	if !monkeyTrouble(false, false) {
		t.Fatal("expected we have trouble but actually not")
	}
}

func TestOneOfThemSmile(t *testing.T) {
	if monkeyTrouble(true, false) {
		t.Fatal("expected we do not have trouble but actually does")
	}
}
