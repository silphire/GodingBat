package ap1

import (
	"strings"
	"testing"
)

func TestWordsWithoutList(t *testing.T) {
	expected := []string{"bb", "ccc"}
	actual := wordsWithoutList([]string{"a", "bb", "b", "ccc"}, 1)
	if len(expected) != len(actual) {
		t.Fatal("length mismatch")
	}
	for i := 0; i < len(expected); i++ {
		if strings.Compare(expected[i], actual[i]) != 0 {
			t.Fatalf("expected \"%s\" but actual is \"%s\" at %d", expected[i], actual[i], i)
		}
	}

	expected = []string{"a", "bb", "b"}
	actual = wordsWithoutList([]string{"a", "bb", "b", "ccc"}, 3)
	if len(expected) != len(actual) {
		t.Fatal("length mismatch")
	}
	for i := 0; i < len(expected); i++ {
		if strings.Compare(expected[i], actual[i]) != 0 {
			t.Fatalf("expected \"%s\" but actual is \"%s\" at %d", expected[i], actual[i], i)
		}
	}

	expected = []string{"a", "bb", "b", "ccc"}
	actual = wordsWithoutList([]string{"a", "bb", "b", "ccc"}, 4)
	if len(expected) != len(actual) {
		t.Fatal("length mismatch")
	}
	for i := 0; i < len(expected); i++ {
		if strings.Compare(expected[i], actual[i]) != 0 {
			t.Fatalf("expected \"%s\" but actual is \"%s\" at %d", expected[i], actual[i], i)
		}
	}
}
