package logic2

import "testing"

func TestMakeBricks(t *testing.T) {
	if !makeBricks(3, 1, 8) {
		t.Fatal("expected true but actually not");
	}
	if makeBricks(3, 1, 9) {
		t.Fatal("expected false but actually not");
	}
	if !makeBricks(3, 2, 10) {
		t.Fatal("expected true but actually not");
	}
}