package array1

import "testing"

func TestBiggerTwo1(t *testing.T) {
	expected := []int{3, 4}
	actual := biggerTwo([]int{1, 2}, []int{3, 4})
	if len(expected) != len(actual) {
		t.Fatal("size mismatch")
	}
	for i := 0; i < 2; i++ {
		if expected[i] != actual[i] {
			t.Fatal("value mismatch")
		}
	}
}

func TestBiggerTwo2(t *testing.T) {
	expected := []int{3, 4}
	actual := biggerTwo([]int{3, 4}, []int{1, 2})
	if len(expected) != len(actual) {
		t.Fatal("size mismatch")
	}
	for i := 0; i < 2; i++ {
		if expected[i] != actual[i] {
			t.Fatal("value mismatch")
		}
	}
}