package warmup2

import "testing"

func TestArrayOneNine(t *testing.T) {
	expected := 1
	actual := arrayCount9([]int{1, 2, 9})
	if expected != actual {
		t.Fatalf("expected \"%d\", but actually \"%d\"", expected, actual)
	}
}

func TestArrayTwoNine(t *testing.T) {
	expected := 2
	actual := arrayCount9([]int{1, 9, 2, 9})
	if expected != actual {
		t.Fatalf("expected \"%d\", but actually \"%d\"", expected, actual)
	}
}
