package string2

import "testing"

func TestEndOther(t *testing.T) {
	if !endOther("Hiabc", "abc") {
		t.Fatal("expected true but actually not")
	}
	if !endOther("AbC", "HiaBc") {
		t.Fatal("expected true but actually not")
	}
	if !endOther("abc", "abXabc") {
		t.Fatal("expected true but actually not")
	}
}
