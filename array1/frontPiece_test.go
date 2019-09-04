package array1

import "testing"

func TestFrontPiece(t *testing.T) {
	expected := []int{1, 2}
	actual := frontPiece([]int{1, 2, 3, 4})
	if len(actual) != 2 {
		t.Fatal("length mismatch")
	}
	for i := 0; i < 2; i++ {
		if expected[i] != actual[i] {
			t.Fatal("content mismatch")
		}
	}
}