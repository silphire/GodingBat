package warmup1

import "testing"

// https://codingbat.com/prob/p101887

func TestAIsLargest(t *testing.T) {
	actual := intMax(3, 2, 1)
	expected := 3

	if actual != expected {
		t.Fatal("expected 3 but actually not")
	}
}
func TestBIsLargest(t *testing.T) {
	actual := intMax(1, 3, 2)
	expected := 3

	if actual != expected {
		t.Fatal("expected 3 but actually not")
	}
}
func TestCIsLargest(t *testing.T) {
	actual := intMax(1, 2, 3)
	expected := 3

	if actual != expected {
		t.Fatal("expected 3 but actually not")
	}
}
