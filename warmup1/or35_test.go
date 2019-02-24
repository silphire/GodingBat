package warmup1

import "testing"

func TestMultiple3(t *testing.T) {
	actual := or35(3)
	if !actual {
		t.Fatal("expected true but actually not")
	}
}
func TestMultiple5(t *testing.T) {
	actual := or35(15)
	if !actual {
		t.Fatal("expected true but actually not")
	}
}
func TestNotMultiple35(t *testing.T) {
	actual := or35(19)
	if actual {
		t.Fatal("expected false but actually not")
	}
}
