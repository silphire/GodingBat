package array3

import "testing"

func TestMaxMirror(t *testing.T) {
	actual := 0

	actual = maxMirror([]int{1, 2, 3, 8, 9, 3, 2, 1})
	if actual != 3 {
		t.Fatalf("expected 3 but actual is %d", actual)
	}

	actual = maxMirror([]int{1, 2, 1, 4})
	if actual != 3 {
		t.Fatalf("expected 3 but actual is %d", actual)
	}

	actual = maxMirror([]int{7, 1, 2, 9, 7, 2, 1})
	if actual != 2 {
		t.Fatalf("expected 2 but actual is %d", actual)
	}
}
