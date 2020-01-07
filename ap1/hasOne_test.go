package ap1

import "testing"

func TestHasOne(t *testing.T) {
	if !hasOne(10) {
		t.Fatal("expected true but actual is not")
	}
	if hasOne(22) {
		t.Fatal("expected false but actual is not")
	}
	if hasOne(220) {
		t.Fatal("expected false but actual is not")
	}
}
