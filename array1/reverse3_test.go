package array1

import "testing"
import "reflect"

func TestReverse3(t *testing.T) {
	expected := []int{3, 2, 1}
	actual := reverse3([]int {1, 2, 3})
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("not expected result")
	}
}