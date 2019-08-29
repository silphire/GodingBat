package array1

import "testing"

func TestMakePi(t *testing.T) {
	expected := []int{3, 1, 4}
	actual := makePi()

	if len(expected) != len(actual) {
		t.Fatal("length mismatch")
	}
	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Fatal("content mismatch")
		}
	}
}