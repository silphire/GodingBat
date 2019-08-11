package array1

import "testing"

func TestMiddleWay(t *testing.T) {
	expected := []int{2, 5}
	actual := middleWay([]int{1, 2, 3}, []int{4, 5, 6})
	if !(actual[0] == expected[0] && actual[1] == expected[1]) {
		t.Fatal("expected true but not")
	}
}