package array1

import "testing"

func TestMakeMiddle1(t *testing.T) {
	expected := []int{2, 3}
	actual := makeMiddle([]int{1, 2, 3, 4})
	if expected[0] != actual[0] || expected[1] != actual[1] {
		t.Fatal("wrong result")
	}
}