package ap1

import (
	"strings"
	"testing"
)

func TestWordsFront(t *testing.T) {
	expected := []string{"a"}
	actual := wordsFront([]string{"a", "b", "c", "d"}, 1)
	if len(expected) != len(actual) {
		t.Fatal("length mismatch")
	}
	for i := 0; i < len(expected); i++ {
		if strings.Compare(expected[i], actual[i]) != 0 {
			t.Fatalf("expected \"%s\" but actual is \"%s\" at %d", expected[i], actual[i], i)
		}
	}

	expected = []string{"a", "b"}
	actual = wordsFront([]string{"a", "b", "c", "d"}, 2)
	if len(expected) != len(actual) {
		t.Fatal("length mismatch")
	}
	for i := 0; i < len(expected); i++ {
		if strings.Compare(expected[i], actual[i]) != 0 {
			t.Fatalf("expected \"%s\" but actual is \"%s\" at %d", expected[i], actual[i], i)
		}
	}

	expected = []string{"a", "b", "c"}
	actual = wordsFront([]string{"a", "b", "c", "d"}, 3)
	if len(expected) != len(actual) {
		t.Fatal("length mismatch")
	}
	for i := 0; i < len(expected); i++ {
		if strings.Compare(expected[i], actual[i]) != 0 {
			t.Fatalf("expected \"%s\" but actual is \"%s\" at %d", expected[i], actual[i], i)
		}
	}}
