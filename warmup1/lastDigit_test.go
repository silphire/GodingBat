package warmup1

import "testing"

func TestSameDigit(t *testing.T) {
	actual := lastDigit(25, 15)
	if !actual {
		t.Fatal("expected true but actually not")

	}
}

func TestDifferentDigit(t *testing.T) {
	actual := lastDigit(27, 36)

	if actual {
		t.Fatal("expected false but actually not")
	}
}
