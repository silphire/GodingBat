package array1

import "testing"

func TestSum2(t *testing.T) {
	expected := 3
	actual := sum2([]int{1, 2, 3})
	if expected != actual {
		t.Fatal("value mismatch")
	}
}