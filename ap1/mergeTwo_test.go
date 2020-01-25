package ap1

import (
	"strings"
	"testing"
)

func TestMergeTwo(t *testing.T) {
	expected := []string{"a", "b", "c"}
	actual := mergeTwo([]string{"a", "c", "z"}, []string{"b", "f", "z"}, 3)
	if len(expected) != len(actual) {
		t.Fatal("length mismatch")
	}
	for i := 0; i < len(expected); i++ {
		if strings.Compare(expected[i], actual[i]) != 0 {
			t.Fatalf("expected \"%s\" but actual is \"%s\" at %d", expected[i], actual[i], i)
		}
	}

	expected = []string{"a", "c", "f"}
	actual = mergeTwo([]string{"a", "c", "z"}, []string{"c", "f", "z"}, 3)
	if len(expected) != len(actual) {
		t.Fatal("length mismatch")
	}
	for i := 0; i < len(expected); i++ {
		if strings.Compare(expected[i], actual[i]) != 0 {
			t.Fatalf("expected \"%s\" but actual is \"%s\" at %d", expected[i], actual[i], i)
		}
	}

	expected = []string{"c", "f", "g"}
	actual = mergeTwo([]string{"f", "g", "z"}, []string{"c", "f", "g"}, 3)
	if len(expected) != len(actual) {
		t.Fatal("length mismatch")
	}
	for i := 0; i < len(expected); i++ {
		if strings.Compare(expected[i], actual[i]) != 0 {
			t.Fatalf("expected \"%s\" but actual is \"%s\" at %d", expected[i], actual[i], i)
		}
	}
}
