package array1

import "testing"

func TestMaxEnd31(t *testing.T) {
	actual := maxEnd3([]int{1, 2, 3})
	for i := 0; i < 3; i++ {
		if actual[i] != 3 {
			t.Fatalf("expected %d but actual is %d at %d", 3, actual[i], i)
		}
	}
}

func TestMaxEnd32(t *testing.T) {
	actual := maxEnd3([]int{4, 2, 3})
	for i := 0; i < 3; i++ {
		if actual[i] != 4 {
			t.Fatalf("expected %d but actual is %d at %d", 4, actual[i], i)
		}
	}
}