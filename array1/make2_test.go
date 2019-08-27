package array1

import "testing"

func TestMake21(t *testing.T) {
	expected := []int{1, 2}
	actual := make2([]int{1, 2}, []int{3, 4})

	if len(expected) != len(actual) {
		t.Fatal("length mismatch")
	}
	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Fatal("contents mismatch")
		}
	}
}

func TestMake22(t *testing.T) {
	expected := []int{1, 3}
	actual := make2([]int{1}, []int{3, 4})
	
	if len(expected) != len(actual) {
		t.Fatal("length mismatch")
	}
	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Fatal("contents mismatch")
		}
	}
}

func TestMake23(t *testing.T) {
	expected := []int{3, 4}
	actual := make2([]int{}, []int{3, 4})
	
	if len(expected) != len(actual) {
		t.Fatal("length mismatch")
	}
	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Fatal("contents mismatch")
		}
	}
}