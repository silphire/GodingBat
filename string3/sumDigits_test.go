package string3

import "testing"

func TestSumDigits(t *testing.T) {
	actual := sumDigits("aa1bc2d3")
	if actual != 6 {
		t.Fatalf("expected 6 but actual is %d", actual)
	}

	actual = sumDigits("aa11b33")
	if actual != 8 {
		t.Fatalf("expected 8 but actual is %d", actual)
	}

	actual = sumDigits("Chocolate")
	if actual != 0 {
		t.Fatalf("expected 0 but actual is %d", actual)
	}
}
