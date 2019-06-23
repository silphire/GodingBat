package warmup2

import "testing"

func TestNoTriples(t *testing.T) {
	if noTriples([]int{1, 1, 2, 2, 1}) {
		t.Fatal("expected false but actually not")
	}
}

func TestHaveTriples(t *testing.T) {
	if !noTriples([]int{1, 1, 2, 2, 2}) {
		t.Fatal("expected false but actually not")
	}
}
