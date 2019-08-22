package array1

import "testing"

func TestMakeEnds(t *testing.T) {
	expected := []int{1, 3}
	actual := makeEnds([]int{1, 2, 3})
	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Fatal("mismatched with expected")
		}
	}
}