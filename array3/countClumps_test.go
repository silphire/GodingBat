package array3

import "testing"

func TestCountClumps(t *testing.T) {
	actual := 0
	actual = countClumps([]int{1, 2, 2, 3, 4, 4})
	if actual != 2 {
		t.Fatalf("expected 2 but actual is %d", actual)
	}
	actual = countClumps([]int{1, 1, 2, 1, 1})
	if actual != 2 {
		t.Fatalf("expected 2 but actual is %d", actual)
	}
	actual = countClumps([]int{1, 1, 1, 1, 1})
	if actual != 1 {
		t.Fatalf("expected 1 but actual is %d", actual)
	}
}
