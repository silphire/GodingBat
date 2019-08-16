package array1

import "testing"

func TestMidThree(t *testing.T) {
	expected := []int{2, 3, 4}
	actual := midThree([]int{1, 2, 3, 4, 5})
	if len(expected) != len(actual) {
		t.Fatal("length different")
	}
	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Fatal("wrong result")
		}
	}
}