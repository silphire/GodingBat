package ap1

import "testing"

func TestDividesSelf(t *testing.T) {
	if !dividesSelf(128) {
		t.Fatal("expected true but actual is not")
	}
	if !dividesSelf(12) {
		t.Fatal("expected true but actual is not")
	}
	if dividesSelf(120) {
		t.Fatal("expected false but actual is not")
	}
	if dividesSelf(123) {
		t.Fatal("expected false but actual is not")
	}
}
