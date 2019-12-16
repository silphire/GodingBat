package array2

import (
	"strings"
	"testing"
)

func TestFizzArray2(t *testing.T) {
	expected := []string{"0", "1", "2", "3"}
	actual := fizzArray2(4)
	if len(expected) != len(actual) {
		t.Fatal("length mismatch")
	}
	for i := 0; i < len(expected); i++ {
		if strings.Compare(expected[i], actual[i]) != 0 {
			t.Fatalf("expected \"%s\" but actual is \"%s\" at %d", expected[i], actual[i], i)
		}
	}

	expected = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	actual = fizzArray2(10)
	if len(expected) != len(actual) {
		t.Fatal("length mismatch")
	}
	for i := 0; i < len(expected); i++ {
		if strings.Compare(expected[i], actual[i]) != 0 {
			t.Fatalf("expected \"%s\" but actual is \"%s\" at %d", expected[i], actual[i], i)
		}
	}

	expected = []string{"0", "1"}
	actual = fizzArray2(2)
	if len(expected) != len(actual) {
		t.Fatal("length mismatch")
	}
	for i := 0; i < len(expected); i++ {
		if strings.Compare(expected[i], actual[i]) != 0 {
			t.Fatalf("expected \"%s\" but actual is \"%s\" at %d", expected[i], actual[i], i)
		}
	}
}
