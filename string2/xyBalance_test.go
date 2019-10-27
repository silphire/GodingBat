package string2

import "testing"

func TestXyBalance(t *testing.T) {
	if !xyBalance("aaxbby") {
		t.Fatal("expected true but actually not")
	}
	if xyBalance("aaxbb") {
		t.Fatal("expected false but actually not")
	}
	if xyBalance("yaaxbb") {
		t.Fatal("expected false but actually not")
	}
}
