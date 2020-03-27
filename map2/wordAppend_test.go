package map2

import (
	"strings"
	"testing"
)

func TestWordAppend(t *testing.T) {
	actual := wordAppend([]string{"a", "b", "a"})
	expected := "a"
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}

	actual = wordAppend([]string{"a", "b", "a", "c", "a", "d", "a"})
	expected = "aa"
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}

	actual = wordAppend([]string{"a", "", "a"})
	expected = "a"
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\" but actual is \"%s\"", expected, actual)
	}
}
