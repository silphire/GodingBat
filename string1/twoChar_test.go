package string1

import "testing"
import "strings"

func TestTwoChar1(t *testing.T) {
	expected := "he"
	actual := twoChar("hello", 0)
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\", but actual is \"%s\"", expected, actual)
	}
}

func TestTwoChar2(t *testing.T) {
	expected := "lo"
	actual := twoChar("hello", 3)
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\", but actual is \"%s\"", expected, actual)
	}
}

func TestTwoChar3(t *testing.T) {
	expected := "he"
	actual := twoChar("hello", 4)
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\", but actual is \"%s\"", expected, actual)
	}
}