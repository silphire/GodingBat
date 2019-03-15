package warmup1

import "testing"

func TestFirstIsLargerAndInRange(t *testing.T) {
	actual := max1020(19, 11)
	expected := 19
	if actual != expected {
		t.Fatalf("expected %d, but actually %d", expected, actual)
	}
}

func TestSecondIsLargerAndInRange(t *testing.T) {
	actual := max1020(14, 18)
	expected := 18
	if actual != expected {
		t.Fatalf("expected %d, but actually %d", expected, actual)
	}
}

func TestFirstIsLargerAndNotInRange(t *testing.T) {
	actual := max1020(66, 19)
	expected := 19
	if actual != expected {
		t.Fatalf("expected %d, but actually %d", expected, actual)
	}
}

func TestSecondIsLargerAndNotInRange(t *testing.T) {
	actual := max1020(14, 21)
	expected := 14
	if actual != expected {
		t.Fatalf("expected %d, but actually %d", expected, actual)
	}
}
func TestBothAreNotInRange(t *testing.T) {
	actual := max1020(3, 21)
	expected := 0
	if actual != expected {
		t.Fatalf("expected %d, but actually %d", expected, actual)
	}
}
