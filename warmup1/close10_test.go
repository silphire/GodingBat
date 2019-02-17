package warmup1

import "testing"

func TestFirstIsNearest(t *testing.T) {
	if close10(9, 12) != 9 {
		t.Fatal("expected first is nearest, but actually not.")
	}
}

func TestSecondIsNearest(t *testing.T) {
	if close10(12, 9) != 9 {
		t.Fatal("expected second is nearest, but actually not")
	}
}

func TestTie(t *testing.T) {
	if close10(8, 12) != 0 {
		t.Fatal("expected tie, but actually not")
	}
}
