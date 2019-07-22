package string1

import "testing"
import "strings"

func TestConCat(t *testing.T) {
	expected := "concat"
	actual := conCat("con", "cat")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\", but actual is \"%s\"", expected, actual)
	}
}

func TestConCatWithEmpty(t *testing.T) {
	expected := "con"
	actual := conCat("con", "")
	if strings.Compare(expected, actual) != 0 {
		t.Fatalf("expected \"%s\", but actual is \"%s\"", expected, actual)
	}
}