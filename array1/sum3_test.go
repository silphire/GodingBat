package array1

import "testing"

func TestSum3(t *testing.T) {
	expected := 6
	actual := sum3([]int{1, 2, 3})
	if expected != actual {
		t.Fatalf("expected %d but actual is %d", expected, actual)
	}
}