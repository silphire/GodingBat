package warmup1

import "testing"

func TestDifferentValue(t *testing.T) {
	actual := sumDouble(5, 7)
	expected := 12

	if actual != expected {
		t.Fatalf("expected %d, but %d actually", expected, actual)
	}
}

func TestSameValue(t *testing.T) {
	actual := sumDouble(4, 4)
	expected := 64

	if actual != expected {
		t.Fatalf("expected %d, but %d actually", expected, actual)
	}
}
