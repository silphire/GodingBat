package string1

import "strings"
import "testing"

func TestMakeAbba(t *testing.T) {
	expected := "abba"
	actual := makeAbba("a", "b")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\", actual is \"%s\"", expected, actual)
	}
}