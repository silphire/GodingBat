package ap1

import (
	"strings"
	"testing"
)

func TestWordsWithout(t *testing.T) {
	expected := []string{"b", "c"}
	actual := wordsWithout([]string{"a", "b", "c", "a"}, "a")
	if len(expected) != len(actual) {
		t.Fatal("length mismatch")
	}
	for i := 0; i < len(expected); i++ {
		if strings.Compare(expected[i], actual[i]) != 0 {
			t.Fatalf("expected \"%s\" but actual is \"%s\" at %d", expected[i], actual[i], i)
		}
	}

	expected = []string{"a", "c", "a"}
	actual = wordsWithout([]string{"a", "b", "c", "a"}, "b")
	if len(expected) != len(actual) {
		t.Fatal("length mismatch")
	}
	for i := 0; i < len(expected); i++ {
		if strings.Compare(expected[i], actual[i]) != 0 {
			t.Fatalf("expected \"%s\" but actual is \"%s\" at %d", expected[i], actual[i], i)
		}
	}

	expected = []string{"a", "b", "a"}
	actual = wordsWithout([]string{"a", "b", "c", "a"}, "c")
	if len(expected) != len(actual) {
		t.Fatal("length mismatch")
	}
	for i := 0; i < len(expected); i++ {
		if strings.Compare(expected[i], actual[i]) != 0 {
			t.Fatalf("expected \"%s\" but actual is \"%s\" at %d", expected[i], actual[i], i)
		}
	}
}
