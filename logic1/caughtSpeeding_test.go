package logic1

import "testing"

func TestCaughtSpeeding(t *testing.T) {
	if caughtSpeeding(60, false) != 0 {
		t.Fatal("expected 0 but actually not")
	}
	if caughtSpeeding(65, false) != 1 {
		t.Fatal("expected 1 but actually not")
	}
	if caughtSpeeding(65, true) != 0 {
		t.Fatal("expected 0 but actually not")
	}
}