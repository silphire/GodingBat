package array1

import "testing"

func TestFix231(t *testing.T) {
	expected := []int{1, 2, 0}
	actual := fix23([]int{1, 2, 3})

	if len(expected) != len(actual) {
		t.Fatal("length different between expected and actual")
	}
	for i := 0; i < len(actual); i++ {
		if expected[i] != actual[i] {
			t.Fatalf("elements at the index %d are different", i)
		}
	}
}

func TestFix232(t *testing.T) {
	expected := []int{2, 0, 5}
	actual := fix23([]int{2, 3, 5})

	if len(expected) != len(actual) {
		t.Fatal("length different between expected and actual")
	}
	for i := 0; i < len(actual); i++ {
		if expected[i] != actual[i] {
			t.Fatalf("elements at the index %d are different", i)
		}
	}
}

func TestFix233(t *testing.T) {
	expected := []int{3, 2, 1}
	actual := fix23([]int{3, 2, 1})

	if len(expected) != len(actual) {
		t.Fatal("length different between expected and actual")
	}
	for i := 0; i < len(actual); i++ {
		if expected[i] != actual[i] {
			t.Fatalf("elements at the index %d are different", i)
		}
	}
}